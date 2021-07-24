package main

import (
	"reflect"
	"testing"
)

var s0 = &State { "First" }
var s1 = &State { "Second" }
var s2 = &State { "Third" }
var s3 = &State { "Fourth" }
var inputRuneA = 'a'
var inputRuneB = 'b'

func Test_NFARuleBook(t *testing.T) {
	rb := NFARuleBook {
		Rules: []NFARule{
			{ TargetState: s0, Char: &inputRuneA, NextState: s0 },
			{ TargetState: s0, Char: &inputRuneB, NextState: s0 },
			{ TargetState: s0, Char: &inputRuneB, NextState: s1 },
			{ TargetState: s1, Char: &inputRuneA, NextState: s2 },
			{ TargetState: s1, Char: &inputRuneB, NextState: s2 },
			{ TargetState: s2, Char: &inputRuneA, NextState: s3 },
			{ TargetState: s2, Char: &inputRuneB, NextState: s3 },
		},
	}

	tests := []struct{
		StartStates []*State
		Input *rune
		Tgt []*State
	}{
		{ []*State{s0}, &inputRuneB, []*State{s0, s1} },
		{ []*State{s0}, &inputRuneA, []*State{s0} },
		{ []*State{s0, s1}, &inputRuneA, []*State{s0, s2} },
		{ []*State{s0, s2}, &inputRuneA, []*State{s0, s3} },
	}

	for _, tt := range tests {
		newStates := rb.AdvanceStates(tt.Input, tt.StartStates)

		if !reflect.DeepEqual(tt.Tgt, newStates) {
			t.Errorf("For start states %v, input '%c': Wanted %v, got %v",
				tt.StartStates,
				*tt.Input,
				tt.Tgt,
				newStates,
			)
		}

	}

}

func Test_NFA(t *testing.T) {
	rb := NFARuleBook {
		Rules: []NFARule{
			{ TargetState: s0, Char: &inputRuneA, NextState: s0 },
			{ TargetState: s0, Char: &inputRuneB, NextState: s0 },
			{ TargetState: s0, Char: &inputRuneB, NextState: s1 },
			{ TargetState: s1, Char: &inputRuneA, NextState: s2 },
			{ TargetState: s1, Char: &inputRuneB, NextState: s2 },
			{ TargetState: s2, Char: &inputRuneA, NextState: s3 },
			{ TargetState: s2, Char: &inputRuneB, NextState: s3 },
		},
	}

	tests := []struct{
		StartStates []*State
		AcceptStates []*State
		Input *rune
		ShouldAccept bool
	}{
		{ []*State{s0}, []*State{s1}, &inputRuneA, false },
		{ []*State{s0}, []*State{s1}, &inputRuneB, true },
		{ []*State{s0}, []*State{s0}, &inputRuneB, true },
		{ []*State{s0}, []*State{s1}, &inputRuneA, false },
		{ []*State{s0, s1}, []*State{s3}, &inputRuneA, false },
		{ []*State{s0, s2}, []*State{s0, s3}, &inputRuneA, true},
		{ []*State{s2}, []*State{s3}, &inputRuneA, true},
	}

	for _, tt := range tests {
		nfa := NewNFA(rb, tt.StartStates, tt.AcceptStates)
		nfa.ProcessRune(tt.Input)
		if nfa.Accepting() != tt.ShouldAccept {
			t.Errorf("For start states %v, accept states %v, input '%c': Wanted %v, got %v",
				tt.StartStates,
				tt.AcceptStates,
				*tt.Input,
				tt.ShouldAccept,
				nfa.Accepting(),
			)
		}

	}
}

func Test_NFARBFreeMoves(t *testing.T) {
	var s0 = &State { "Start" }
	var s1 = &State { "Mult2_0" }
	// var s2 = &State { "Mult2_1" }

	var s3 = &State { "Mult3_0" }
	//var s4 = &State { "Mult3_1" }
	// var s5 = &State { "Mult3_2" }

	rb := NFARuleBook {
		Rules: []NFARule{
			{ TargetState: s0, Char: nil, NextState: s1 },
			{ TargetState: s0, Char: nil, NextState: s3 },
			/*
			{ TargetState: s1, Char: &inputRuneA, NextState: s2 },
			{ TargetState: s2, Char: &inputRuneA, NextState: s1 },
			{ TargetState: s3, Char: &inputRuneA, NextState: s4 },
			{ TargetState: s4, Char: &inputRuneA, NextState: s5 },
			{ TargetState: s5, Char: &inputRuneA, NextState: s3 },
			*/
		},
	}

	tests := []struct{
		StartStates []*State
		Input *rune
		Tgt []*State
	}{
		{ []*State{s0}, nil, []*State{s1, s3} },
	}

	for _, tt := range tests {
		newStates := rb.AdvanceStates(tt.Input, tt.StartStates)

		if !reflect.DeepEqual(tt.Tgt, newStates) {
			t.Errorf("For start states %v, input '%c': Wanted %v, got %v",
				tt.StartStates,
				*tt.Input,
				tt.Tgt,
				newStates,
			)
		}

	}
}

func Test_NFARBRecurseFreeMoves(t *testing.T) {
	var stStart = &State { "Start" }
	var stFree0 = &State { "Free0" }
	var stFree1 = &State { "Free1" }
	var stNotFree0 = &State { "NotFree0" }

	rb := NFARuleBook {
		Rules: []NFARule{
			{ TargetState: stStart, Char: nil, NextState: stFree0 },
			{ TargetState: stStart, Char: nil, NextState: stFree1 },
			{ TargetState: stStart, Char: &inputRuneA, NextState: stNotFree0 },
		},
	}

	tests := []struct{
		StartStates []*State
		Tgt []*State
	}{
		{ []*State{stStart}, []*State{stStart, stFree0, stFree1} },
	}

	for _, tt := range tests {
		newStates := rb.FollowFreeMoves(tt.StartStates)

		if !reflect.DeepEqual(tt.Tgt, newStates) {
			t.Errorf("For start states %v, nil input: Wanted %v, got %v",
				tt.StartStates,
				tt.Tgt,
				newStates,
			)
		}

	}
}

func Test_NFAWithFreeMoves(t *testing.T) {
	// Use the rules from the book
	// and try processing some of those strings
	var stStart = &State { "Start" }
	var stBranch0_0 = &State { "Branch0_0" }
	var stBranch0_1 = &State { "Branch0_1" }

	var stBranch1_0 = &State { "Branch1_0" }
	var stBranch1_1 = &State { "Branch1_1" }
	var stBranch1_2 = &State { "Branch1_2" }

	rb := NFARuleBook {
		Rules: []NFARule{
			// Free moves
			{ TargetState: stStart, Char: nil, NextState: stBranch0_0 },
			{ TargetState: stStart, Char: nil, NextState: stBranch1_0 },
			// First branch / loop
			{ TargetState: stBranch0_0, Char: &inputRuneA, NextState: stBranch0_1 },
			{ TargetState: stBranch0_1, Char: &inputRuneA, NextState: stBranch0_0 },

			// Second branch / loop
			{ TargetState: stBranch1_0, Char: &inputRuneA, NextState: stBranch1_1 },
			{ TargetState: stBranch1_1, Char: &inputRuneA, NextState: stBranch1_2 },
			{ TargetState: stBranch1_2, Char: &inputRuneA, NextState: stBranch1_0 },
		},
	}

	tests := []struct{
		Input string
		ShouldAccept bool
	}{
		{  "a", false },
		{  "aa", true },
		{  "aaa", true },
		{  "aaaa", true },
		{  "aaaaa", false },
		{  "aaaaaa", true },
	}

	startStates := []*State{stStart}
	acceptStates := []*State{stBranch0_0, stBranch1_0}
	for _, tt := range tests {
		nfa := NewNFA(rb, startStates, acceptStates)
		nfa.ProcessString(tt.Input)

		if nfa.Accepting() != tt.ShouldAccept {
			t.Errorf("For start states %v, accept states %v, input '%s': Wanted %v, got %v",
				startStates,
				acceptStates,
				tt.Input,
				tt.ShouldAccept,
				nfa.Accepting(),
			)
		}

	}
}
