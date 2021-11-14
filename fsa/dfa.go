package main

type DFARuleBook struct {
	Rules []Rule
}

func (rb *DFARuleBook) FindRule(c rune, st *State) *Rule {
	for _, rule := range rb.Rules {
		if rule.AppliesTo(c, st) {
			return &rule
		}
	}
	return nil
}

func (rb *DFARuleBook) NextState(c rune, st *State) *State {
	rule := rb.FindRule(c, st)
	if rule != nil {
		return rule.NextState
	}

	return nil
}

type DFA struct {
	CurrentState *State
	// Which states are valid as "accept" states
	AcceptStates []*State
	RB           DFARuleBook
}

func (dfa *DFA) Accepting() bool {
	for _, st := range dfa.AcceptStates {
		if st == dfa.CurrentState {
			return true
		}
	}
	return false
}

func (dfa *DFA) ProcessRune(r rune) {
	// Find a rule for this rune
	rule := dfa.RB.FindRule(r, dfa.CurrentState)
	// If a rule is found, update the current state based on that rule
	if rule != nil {
		dfa.CurrentState = rule.Follow()
	}
}

func (dfa *DFA) ProcessString(s string) {
	for _, r := range s {
		dfa.ProcessRune(r)
	}
}

func NewDFA(startState *State, acceptStates []*State, rb DFARuleBook) *DFA {
	return &DFA{
		RB:           rb,
		CurrentState: startState,

		AcceptStates: acceptStates,
	}
}
