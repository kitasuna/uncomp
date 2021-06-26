package simplex

import (
	"fmt"
)

type Machine struct {
	Expr Expr
	env  map[string]Expr
}

func (m *Machine) Step() {
	m.Expr = m.Expr.Reduce(m.env)
}

func (m *Machine) Run(debug bool) {
	if debug {
		fmt.Printf("%v, Env: %v\n", m.Expr, m.env)
	}
	for m.Expr.Reducible() {
		m.Step()
		if debug {
			fmt.Printf("%v, Env: %v\n", m.Expr, m.env)
		}
	}
}

type StmtMachine struct {
	S   Stmt
	Env map[string]Expr
}

func (m *StmtMachine) Step() {
	newStmt, newEnv := m.S.Reduce(m.Env)
	m.S = newStmt
	m.Env = newEnv
}

func (m *StmtMachine) Run(debug bool) {
	if debug {
		fmt.Printf("%v, Env: %v\n", m.S, m.Env)
	}
	for m.S.Reducible() {
		m.Step()
		if debug {
			fmt.Printf("%v, Env: %v\n", m.S, m.Env)
		}
	}
}
