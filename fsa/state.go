package main

type State struct {
	Label string
}

func (st *State) String() string {
	return st.Label
}
