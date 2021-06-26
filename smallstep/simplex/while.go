package simplex

import (
	"fmt"
)

type While struct {
	Condition Expr
	Body      Stmt
}

func (w While) String() string {
	return fmt.Sprintf("while (%v) { %v }", w.Condition, w.Body)
}

func (w While) Reducible() bool {
	return true
}

func (w While) Equals(otherStmt Stmt) bool {
	return false
}

func (w While) Reduce(env map[string]Expr) (Stmt, map[string]Expr) {
	return Conditional{Condition: w.Condition, Consequence: Sequence{First: w.Body, Second: w}, Alternative: Noop{}}, env
}
