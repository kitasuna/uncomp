package main

import (
	"fmt"
)

func main() {
	s0 := &State{ "First" }
	s1 := &State{ "Second" }
	s2 := &State{ "Final" }
	rb := DFARuleBook {
		[]Rule{
			{ TargetState: s0, Char: 'a', NextState: s0 },
			{ TargetState: s0, Char: 'b', NextState: s1 },
			{ TargetState: s1, Char: 'a', NextState: s2 },
			{ TargetState: s1, Char: 'b', NextState: s1 },
			{ TargetState: s2, Char: 'a', NextState: s2 },
			{ TargetState: s2, Char: 'b', NextState: s2 },
		},
	}
	dfa := NewDFA(s0, []*State{ s2 }, rb)

	dfa.ProcessString("abba")

	if dfa.Accepting() {
		fmt.Println("Accepted string!")
	} else {
		fmt.Println("Did not accept string")
	}
}
