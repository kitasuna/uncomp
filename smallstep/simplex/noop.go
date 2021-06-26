package simplex

import (
	"reflect"
)

type Noop struct{}

func (op Noop) String() string {
	return "<no-op>"
}

func (op Noop) Reducible() bool {
	return false
}

func (op Noop) Reduce(env map[string]Expr) (Stmt, map[string]Expr) {
	return op, env
}

func (op Noop) Equals(otherStmt Stmt) bool {
	if reflect.TypeOf(otherStmt).Name() == "Noop" {
		return true
	}

	return false
}
