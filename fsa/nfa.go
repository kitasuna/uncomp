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

func (rb *NFARuleBook) Alphabet() []rune {
	rs := make([]rune, 0)
	exists := make(map[rune]bool)
	for _, rule := range rb.Rules {
		if rule.Char != nil {
			if _, found := exists[*rule.Char]; !found {
				rs = append(rs, *rule.Char)
				exists[*rule.Char] = true
			}
		}
	}

	return rs
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

func SetsAreEqual(s1 []*State, s2 []*State) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, s := range s1 {
		if !StateInSlice(s2, s) {
			return false
		}
	}

	return true
}

type NFA struct {
	RB            NFARuleBook
	CurrentStates []*State
	AcceptStates  []*State
}

func NewNFA(rb NFARuleBook, curr []*State, acc []*State) NFA {
	return NFA{
		RB:            rb,
		CurrentStates: curr,
		AcceptStates:  acc,
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

/*
  As I understand it, we're slowly gonna transform this thing into a DFA sim of the NFA
  So even though the wrapper process seems a little silly, maybe there's some long-term benefit
  in actually doing it...
*/
func (nfa *NFA) SetCurrentStates(ss []*State) {
	nfa.CurrentStates = ss
}

type NFASimulation struct {
	Nfa NFA
}

func NewNFASim(nfa NFA) *NFASimulation {
	return &NFASimulation{
		Nfa: nfa,
	}
}

func (nfas *NFASimulation) NextState(ss []*State, r *rune) []*State {
	nfas.Nfa.SetCurrentStates(ss)
	nfas.Nfa.ProcessRune(r)
	return nfas.Nfa.GetCurrentStates()
}

func (nfas *NFASimulation) RulesFor(ss []*State) []NFARuleDescription {
	alphabet := nfas.Nfa.RB.Alphabet()

	rules := make([]NFARuleDescription, 0)
	for idx, runeElem := range alphabet {
		resultingStates := nfas.NextState(ss, &runeElem)
		rules = append(rules, NFARuleDescription{
			SrcStates:  ss,
			Char:       &(alphabet[idx]),
			DestStates: resultingStates,
		})
	}

	return rules
}
