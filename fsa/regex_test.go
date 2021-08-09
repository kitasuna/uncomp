package main

import (
	// "reflect"
	"fmt"
	"testing"
)

func Test_EmptyRegex(t *testing.T) {
	empty := Empty{}

	nfa := empty.ToNFA()

	nfa.ProcessString("")

	fmt.Printf("CurrentStates: %v\n", nfa.CurrentStates)

	if !nfa.Accepting() {
		t.Error("Expected empty state to accept empty string")
	}
}

func Test_LitRegex_OK(t *testing.T) {
	literalA := Lit{'a'}

	nfa := literalA.ToNFA()

	nfa.ProcessString("a")

	fmt.Printf("CurrentStates: %v\n", nfa.CurrentStates)

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

	concat := Concat { First: literalA, Second: literalB }

	nfa := concat.ToNFA()

	nfa.ProcessString("ab")

	if !nfa.Accepting() {
		t.Error("Expected Concat to accept the string `ab`")
	}
}

func Test_Concat_FailOnIncompleteMatch(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	concat := Concat { First: literalA, Second: literalB }

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

	concat := Concat { First: literalA, Second: Concat { First: literalB, Second: literalC } }

	nfa := concat.ToNFA()

	nfa.ProcessString("abc")

	if !nfa.Accepting() {
		t.Error("Expected Concat to accept the string `abc`")
	}
}

func Test_Choose_FirstOK(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose { First: literalA, Second: literalB }

	nfa := choose.ToNFA()

	nfa.ProcessString("a")

	if !nfa.Accepting() {
		t.Error("Expected Choose to accept the string `a`")
	}
}

func Test_Choose_SecondOK(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose { First: literalA, Second: literalB }

	nfa := choose.ToNFA()

	nfa.ProcessString("b")

	if !nfa.Accepting() {
		t.Error("Expected Choose to accept the string `b`")
	}
}

func Test_Choose_Fail1(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose { First: literalA, Second: literalB }

	nfa := choose.ToNFA()

	nfa.ProcessString("c")

	if nfa.Accepting() {
		t.Error("Expected Choose to not accept the string `c`")
	}
}

func Test_Choose_Fail2(t *testing.T) {
	literalA := Lit{'a'}
	literalB := Lit{'b'}

	choose := Choose { First: literalA, Second: literalB }

	nfa := choose.ToNFA()

	nfa.ProcessString("aa")

	if nfa.Accepting() {
		t.Error("Expected Choose to not accept the string `aa`")
	}
}

func Test_Repeat_EmptyOK(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat { Pattern: literalA }

	nfa := repeat.ToNFA()

	nfa.ProcessString("")

	if !nfa.Accepting() {
		t.Error("Expected Repeat to accept the empty string")
	}
}

func Test_Repeat_SingleOK(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat { Pattern: literalA }

	nfa := repeat.ToNFA()

	nfa.ProcessString("a")

	if !nfa.Accepting() {
		t.Error("Expected Repeat to accept the string `a`")
	}
}

func Test_Repeat_MultipleOK(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat { Pattern: literalA }

	nfa := repeat.ToNFA()

	nfa.ProcessString("aaaa")

	if !nfa.Accepting() {
		t.Error("Expected Repeat to accept the string `a`")
	}
}

func Test_Repeat_Fail1(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat { Pattern: literalA }

	nfa := repeat.ToNFA()

	nfa.ProcessString("b")

	if nfa.Accepting() {
		t.Error("Expected Repeat to not accept the string `b`")
	}
}

func Test_Repeat_Fail2(t *testing.T) {
	literalA := Lit{'a'}

	repeat := Repeat { Pattern: literalA }

	nfa := repeat.ToNFA()

	nfa.ProcessString("aab")

	if nfa.Accepting() {
		t.Error("Expected Repeat to not accept the string `aab`")
	}
}
