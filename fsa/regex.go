package main

import (
	"fmt"
	"strings"
)

type Regex interface {
	GetPrecedence() int
	String() string
	ToNFA() NFA
}

func bracket(pattern Regex, otherPrecedence int) string {
	if pattern.GetPrecedence() < otherPrecedence {
		return "(" + pattern.String() + ")"
	}

	return pattern.String()
}

type Empty struct {}

func (e Empty) String() string {
  return ""
}

func (e Empty) ToNFA() NFA{
	rulebook := NFARuleBook{ Rules: []NFARule{} }
	startState := &State{ Label: "EmptyStart" }
	return NewNFA(rulebook, []*State{startState}, []*State{startState})
}

func (e Empty) GetPrecedence() int {
  return 3
}


type Lit struct {
  r rune
}

func (l Lit) String() string {
  return string(l.r)
}


func (l Lit) GetPrecedence() int {
  return 3
}

func (l Lit) ToNFA() NFA {
	startState := &State { Label: fmt.Sprintf("Start(%v)", string(l.r)) }
	readSingleRuneState := &State { Label: fmt.Sprintf("readSingleRune%v", string(l.r)) }
	rulebook := NFARuleBook{
		Rules: []NFARule{
			{
				TargetState: startState,
				Char: &l.r,
				NextState: readSingleRuneState,
			},
		},
	}

	return NewNFA(rulebook, []*State{startState}, []*State{readSingleRuneState})
}

type Concat struct {
	First Regex
	Second Regex
}

func (c Concat) String() string {
	// Not sure on this implementation...
	return bracket(c.First, c.GetPrecedence()) + bracket(c.Second, c.GetPrecedence())
}

func (c Concat) GetPrecedence() int {
	return 0
}

func (c Concat) ToNFA() NFA {
	nfa1 := c.First.ToNFA()
	nfa2 := c.Second.ToNFA()
	rules1 := nfa1.RB.Rules
	combinedRules := append(rules1, nfa2.RB.Rules...)

	var freeMoves []NFARule
	for _, st := range nfa1.AcceptStates {
		freeMoves = append(freeMoves, NFARule{ TargetState: st, Char: nil, NextState: nfa2.CurrentStates[0] })
	}

	combinedRules = append(combinedRules, freeMoves...)

	return NewNFA(
		NFARuleBook{ Rules: combinedRules },
		[]*State{nfa1.CurrentStates[0]},
		nfa2.AcceptStates,
	)
}

type Choose struct {
	First Regex
	Second Regex
}

func (c Choose) String() string {
	return strings.Join(
		[]string{
			bracket(c.First, c.GetPrecedence()),
			bracket(c.Second, c.GetPrecedence()),
		},
		"|",
	)
}

func (c Choose) GetPrecedence() int {
	return 0
}

func (c Choose) ToNFA() NFA{
	nfa1 := c.First.ToNFA()
	nfa2 := c.Second.ToNFA()
	rules1 := nfa1.RB.Rules
	combinedRules := append(rules1, nfa2.RB.Rules...)

	initState := &State{"StartChoose"}
	combinedRules = append(combinedRules, NFARule{ TargetState: initState, Char: nil, NextState: nfa1.CurrentStates[0] })
	combinedRules = append(combinedRules, NFARule{ TargetState: initState, Char: nil, NextState: nfa2.CurrentStates[0] })

	var acceptStates []*State
	acceptStates = append(acceptStates, nfa1.AcceptStates...)
	acceptStates = append(acceptStates, nfa2.AcceptStates...)

	return NewNFA(
		NFARuleBook{ Rules: combinedRules },
		[]*State{ initState },
		acceptStates,
	)
}

type Repeat struct {
	Pattern Regex
}

func (r Repeat) String() string {
	return bracket(r.Pattern, r.GetPrecedence()) + "*"
}

func (r Repeat) GetPrecedence() int {
	return 2
}

func (r Repeat) ToNFA() NFA{
	nfa := r.Pattern.ToNFA()

	// Adding this so it can accept the empty string
	addlStartState := &State{ "StartAcceptsEmpty" }

	combinedRules := nfa.RB.Rules
	for _, s := range nfa.AcceptStates {
		combinedRules = append(combinedRules,
		NFARule{TargetState: s, Char: nil, NextState:addlStartState},
		)
	}
	combinedRules = append(combinedRules,
		NFARule{TargetState: addlStartState, Char: nil, NextState: nfa.CurrentStates[0]},
	)
	return NewNFA(
		NFARuleBook{ Rules: combinedRules },
		[]*State{addlStartState},
		append(nfa.AcceptStates, addlStartState),
	)
}
