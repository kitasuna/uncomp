package main

type NFARuleBook struct {
	Rules []NFARule
}

func (rb *NFARuleBook) AdvanceStates(c *rune, sts []*State) []*State {
	returnStates := make([]*State, 0)
	for _, st := range sts {
		// Find any rules for this pair
		rules := rb.FindAllRules(c, st)
		for _, rule := range rules {
			newState := rule.Follow()
			if newState != nil && !StateInSlice(returnStates, newState) {
				returnStates = append(returnStates, newState)
			}
		}
	}

	return returnStates
}

func (rb *NFARuleBook) FindAllRules(c *rune, st *State) []NFARule {
	returnRules := make([]NFARule, 0)
	for _, rule := range rb.Rules {
		if rule.AppliesTo(c, st) {
			returnRules = append(returnRules, rule)
		}
	}

	return returnRules
}

func (rb *NFARuleBook) FollowFreeMoves(sts []*State) []*State {
	moreStates := rb.AdvanceStates(nil, sts)

	if IsSubsetOf(sts, moreStates) {
		return sts
	} else {
		return rb.FollowFreeMoves(append(sts, moreStates...))
	}
}

// *sigh* Go. Why do you make me do these things.
func StateInSlice(sts []*State, st *State) bool {
	for _, state := range sts {
		if st == state {
			return true
		}
	}

	return false
}

// *sigh part 2*
func IsSubsetOf(checkSet []*State, maybeSubset []*State) bool {
	for _, s := range maybeSubset {
		if !StateInSlice(checkSet, s) {
			return false
		}
	}

	return true
}

type NFA struct {
	RB NFARuleBook
	CurrentStates []*State
	AcceptStates []*State
}

func NewNFA(rb NFARuleBook, curr []*State, acc []*State) NFA {
	return NFA {
		RB: rb,
		CurrentStates: curr,
		AcceptStates: acc,
	}
}

func (nfa *NFA) GetCurrentStates() []*State {
	newCurrentStates := nfa.RB.FollowFreeMoves(nfa.CurrentStates)
	return newCurrentStates
}

func (nfa *NFA) Accepting() bool {
	for _, curr := range nfa.GetCurrentStates() {
		if StateInSlice(nfa.AcceptStates, curr) {
			return true
		}
	}

	return false
}

func (nfa *NFA) ProcessRune(r *rune) {
	nfa.CurrentStates = nfa.RB.AdvanceStates(r, nfa.GetCurrentStates())
}

func (nfa *NFA) ProcessString(str string) {
	for _, r := range str {
		nfa.ProcessRune(&r)
	}
}
