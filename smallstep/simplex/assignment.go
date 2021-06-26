package simplex

import (
	"fmt"
)

type Assignment struct {
	Name string
	X    Expr
}

func (s Assignment) String() string {
	return fmt.Sprintf("%v = %v", s.Name, s.X)
}

func (s Assignment) Reducible() bool {
	return true
}

func (s Assignment) Equals(_ Stmt) bool {
	return false
}

func (s Assignment) Reduce(env map[string]Expr) (Stmt, map[string]Expr) {
	if s.X.Reducible() {
		return Assignment{Name: s.Name, X: s.X.Reduce(env)}, env
	} else {
		op := Noop{}
		newEnv := updateEnv(env, s.Name, s.X)
		return op, newEnv
	}
}

func updateEnv(m map[string]Expr, keyToUpdate string, newVal Expr) map[string]Expr {
	updated := make(map[string]Expr)
	for k, v := range m {
		if k != keyToUpdate {
			updated[k] = v
		}
	}
	updated[keyToUpdate] = newVal
	return updated
}
