package main

import (
	// "reflect"
	"testing"
)

func Test_EmptyRegex(t *testing.T) {
	empty := Empty{}

	nfa := empty.ToNFA()

	nfa.ProcessString("")

	if !nfa.Accepting() {
		t.Error("Expected empty state to accept empty string")
	}
}

func Test_LitRegex_OK(t *testing.T) {
	literalA := Lit{'b'}

	nfa := literalA.ToNFA()

	nfa.ProcessString("b")

	if !nfa.Accepting() {
		t.Error("Expected Lit to accept single occurrence of its specified rune")
	}
}

func Test_LitRegex_FailOnRepeat(t *testing.T) {
	literalA := Lit{'a'}

	nfa := literalA.ToNFA()

	nfa.ProcessString("aa")

	if nfa.Accepting() {
		t.Error("Expected Lit to NOT accept repeat occurrences of its specified rune")
	}
}

func Test_LitRegex_FailOnDiffRune(t *testing.T) {
	literalA := Lit{'a'}

	nfa := literalA.ToNFA()

	nfa.ProcessString("b")

	if nfa.Accepting() {
		t.Error("Expected Lit to NOT accept something other than its specified rune")
	}
}

func Test_Concat_OK(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	concat := Concat{First: literalA, Second: literalB}

	nfa := concat.ToNFA()

	nfa.ProcessString("ab")

	if !nfa.Accepting() {
		t.Error("Expected Concat to accept the string `ab`")
	}
}

func Test_Concat_FailOnIncompleteMatch(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	concat := Concat{First: literalA, Second: literalB}

	nfa := concat.ToNFA()

	nfa.ProcessString("a")

	if nfa.Accepting() {
		t.Error("Expected Concat to fail on the string `a`")
	}
}

func Test_Concat_Recursion(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}
	literalC := Lit{'c'}

	concat := Concat{First: literalA, Second: Concat{First: literalB, Second: literalC}}

	nfa := concat.ToNFA()

	nfa.ProcessString("abc")

	if !nfa.Accepting() {
		t.Error("Expected Concat to accept the string `abc`")
	}
}

func Test_Concat_Choose(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}
	empty := Empty{}

	regex := Concat{
		First: literalA,
		Second: Choose{
			First:  literalB,
			Second: empty,
		},
	}

	tests := []struct {
		Str     string
		Accepts bool
	}{
		{
			Str:     "a",
			Accepts: true,
		},
		{
			Str:     "ab",
			Accepts: true,
		},
	}

	for _, tt := range tests {
		nfa := regex.ToNFA()
		nfa.ProcessString(tt.Str)

		if nfa.Accepting() != tt.Accepts {
			t.Errorf("Expected NFA to return %v the string `%v`, but got %v", tt.Accepts, tt.Str, nfa.Accepting())
		}
	}
}

func Test_Choose_FirstOK(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose{First: literalA, Second: literalB}

	nfa := choose.ToNFA()

	nfa.ProcessString("a")

	if !nfa.Accepting() {
		t.Error("Expected Choose to accept the string `a`")
	}
}

func Test_Choose_EmptyFirstOK(t *testing.T) {
	literalB := Lit{'b'}

	choose := Choose{First: Empty{}, Second: literalB}

	nfa := choose.ToNFA()

	nfa.ProcessString("")

	if !nfa.Accepting() {
		t.Error("Expected Choose to accept the empty string")
	}
}

func Test_Choose_EmptySecondOK(t *testing.T) {
	literalB := Lit{'b'}

	choose := Choose{First: literalB, Second: Empty{}}

	nfa := choose.ToNFA()
	nfa.ProcessString("b")

	if !nfa.Accepting() {
		t.Error("Expected Choose to accept the empty string")
	}
}

func Test_Choose_SecondOK(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose{First: literalA, Second: literalB}

	nfa := choose.ToNFA()

	nfa.ProcessString("b")

	if !nfa.Accepting() {
		t.Error("Expected Choose to accept the string `b`")
	}
}

func Test_Choose_Fail1(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose{First: literalA, Second: literalB}

	nfa := choose.ToNFA()

	nfa.ProcessString("c")

	if nfa.Accepting() {
		t.Error("Expected Choose to not accept the string `c`")
	}
}

func Test_Choose_Fail2(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose{First: literalA, Second: literalB}

	nfa := choose.ToNFA()

	nfa.ProcessString("aa")

	if nfa.Accepting() {
		t.Error("Expected Choose to not accept the string `aa`")
	}
}

func Test_Repeat_EmptyOK(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat{Pattern: literalA}

	nfa := repeat.ToNFA()

	nfa.ProcessString("")

	if !nfa.Accepting() {
		t.Error("Expected Repeat to accept the empty string")
	}
}

func Test_Repeat_SingleOK(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat{Pattern: literalA}

	nfa := repeat.ToNFA()

	nfa.ProcessString("a")

	if !nfa.Accepting() {
		t.Error("Expected Repeat to accept the string `a`")
	}
}

func Test_Repeat_MultipleOK(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat{Pattern: literalA}

	nfa := repeat.ToNFA()

	nfa.ProcessString("aaaa")

	if !nfa.Accepting() {
		t.Error("Expected Repeat to accept the string `a`")
	}
}

func Test_Repeat_Fail1(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat{Pattern: literalA}

	nfa := repeat.ToNFA()

	nfa.ProcessString("b")

	if nfa.Accepting() {
		t.Error("Expected Repeat to not accept the string `b`")
	}
}

func Test_Repeat_Fail2(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat{Pattern: literalA}

	nfa := repeat.ToNFA()

	nfa.ProcessString("aab")

	if nfa.Accepting() {
		t.Error("Expected Repeat to not accept the string `aab`")
	}
}

func Test_BigOne(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}
	big := Repeat{
		Pattern: Concat{
			First:  literalA,
			Second: Choose{First: Empty{}, Second: literalB},
		},
	}

	tests := []struct {
		Str     string
		Accepts bool
	}{
		{
			Str:     "a",
			Accepts: true,
		},
		{
			Str:     "ab",
			Accepts: true,
		},
		{
			Str:     "ab",
			Accepts: true,
		},
		{
			Str:     "aba",
			Accepts: true,
		},
		{
			Str:     "abab",
			Accepts: true,
		},
		{
			Str:     "abaab",
			Accepts: true,
		},
		{
			Str:     "abba",
			Accepts: false,
		},
	}

	for _, tt := range tests {
		nfa := big.ToNFA()
		nfa.ProcessString(tt.Str)

		if nfa.Accepting() != tt.Accepts {
			t.Errorf("Expected NFA to return %v the string `%v`, but got %v", tt.Accepts, tt.Str, nfa.Accepting())
		}
	}
}
