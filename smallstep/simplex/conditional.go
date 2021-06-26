package simplex

import (
	"fmt"
)

type Conditional struct {
	Condition   Expr
	Consequence Stmt
	Alternative Stmt
}

func (c Conditional) Reducible() bool {
	return true
}

func (c Conditional) String() string {
	return fmt.Sprintf("if ( %v ) { %v } else { %v }", c.Condition, c.Consequence, c.Alternative)
}

func (c Conditional) Equals(other Stmt) bool {
	return false
}

func (c Conditional) Reduce(env map[string]Expr) (Stmt, map[string]Expr) {
	if c.Condition.Reducible() {
		return Conditional{
			Condition:   c.Condition.Reduce(env),
			Consequence: c.Consequence,
			Alternative: c.Alternative,
		}, env
	} else if c.Condition.(Boolean).Equals(Boolean{Val: true}) {
		return c.Consequence, env
	} else {
		return c.Alternative, env
	}
}
