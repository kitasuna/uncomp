package simplex

import (
	"fmt"
)

type Sequence struct {
	First  Stmt
	Second Stmt
}

func (s Sequence) String() string {
	return fmt.Sprintf("%v; %v", s.First, s.Second)
}

func (s Sequence) Reducible() bool {
	return true
}

func (s Sequence) Equals(other Stmt) bool {
	return false
}

func (s Sequence) Reduce(env map[string]Expr) (Stmt, map[string]Expr) {
	if s.First.Equals(Noop{}) {
		return s.Second, env
	} else {
		newFirst, newEnv := s.First.Reduce(env)
		return Sequence{First: newFirst, Second: s.Second}, newEnv
	}
}
