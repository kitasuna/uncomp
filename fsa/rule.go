package main

import (
	"fmt"
)

type Rule struct {
	TargetState *State
	Char rune
	NextState *State
}

func (r *Rule) AppliesTo(c rune, st *State) bool {
	return r.TargetState == st && c == r.Char
}

func (r Rule) String() string {
	return fmt.Sprintf("%v -- %v ---> %v", *r.TargetState, string(r.Char), *r.NextState)
}

func (r *Rule) Follow() *State {
	return r.NextState
}

// Using a separate rule type for NFAs to handle free moves
// without breaking the DFA implementation
type NFARule struct {
	TargetState *State
	Char *rune
	NextState *State
}

func (r *NFARule) AppliesTo(c *rune, st *State) bool {
	if(r.TargetState == st) {
		if c == nil && r.Char == nil {
			return true
		} else if (c == nil) || (r.Char == nil) {
			return false
		} else if *c == *r.Char {
			return true
		}
	}

	return false
}

func (r NFARule) String() string {
	if r.Char != nil {
		return fmt.Sprintf("%v -- %v ---> %v", r.TargetState, string(*r.Char), r.NextState)
	}

	return fmt.Sprintf("%v -- <nil> ---> %v", r.TargetState, r.NextState)
}

func (r *NFARule) Follow() *State {
	return r.NextState
}
