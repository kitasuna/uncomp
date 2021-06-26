package simplex

import (
	"fmt"
)

type Boolean struct {
	Val bool
}

func (b Boolean) Reducible() bool {
	return false
}

func (b Boolean) Reduce(_ map[string]Expr) Expr {
	return b
}

func (b Boolean) String() string {
	if b.Val == true {
		return "True"
	}

	return "False"
}

func (b Boolean) Equals(other Boolean) bool {
	return b.Val == other.Val
}

type LessThanExpr struct {
	Left, Right Expr
}

func (x LessThanExpr) String() string {
	return fmt.Sprintf("%v < %v", x.Left, x.Right)
}

func (x LessThanExpr) Reducible() bool {
	return true
}

func (x LessThanExpr) Reduce(env map[string]Expr) Expr {
	if x.Left.Reducible() {
		return LessThanExpr{Left: x.Left.Reduce(env), Right: x.Right}
	} else if x.Right.Reducible() {
		return LessThanExpr{Left: x.Left, Right: x.Right.Reduce(env)}
	} else if (x.Left).(Number).Val < x.Right.(Number).Val {
		return Boolean{Val: true}
	}

	return Boolean{Val: false}
}
