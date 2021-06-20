package simplex

import (
	"fmt"
)

type Number struct {
	Val int
}

func (n Number) String() string {
	return fmt.Sprintf("%v", n.Val)
}

func (n Number) Reducible() bool {
	return false
}

func (n Number) Reduce(_ map[string]Expr) Expr {
	return n
}

type AddExpr struct {
	Left, Right Expr
}

func (x AddExpr) String() string {
	return fmt.Sprintf("%v + %v", x.Left, x.Right)
}

func (n AddExpr) Reducible() bool {
	return true
}

func (expr AddExpr) Reduce(env map[string]Expr) Expr {
	if expr.Left.Reducible() {
		return AddExpr{Left: expr.Left.Reduce(env), Right: expr.Right}
	} else if expr.Right.Reducible() {
		return AddExpr{Left: expr.Left, Right: expr.Right.Reduce(env)}
	} else {
		return Number{Val: expr.Left.(Number).Val + expr.Right.(Number).Val}
	}
}

type MultExpr struct {
	Left, Right Expr
}

func (x MultExpr) String() string {
	return fmt.Sprintf("%v * %v", x.Left, x.Right)
}

func (n MultExpr) Reducible() bool {
	return true
}

func (expr MultExpr) Reduce(env map[string]Expr) Expr {
	if expr.Left.Reducible() {
		return MultExpr{Left: expr.Left.Reduce(env), Right: expr.Right}
	} else if expr.Right.Reducible() {
		return MultExpr{Left: expr.Left, Right: expr.Right.Reduce(env)}
	} else {
		return Number{Val: expr.Left.(Number).Val * expr.Right.(Number).Val}
	}
}

type SubtrExpr struct {
	Left, Right Expr
}

func (x SubtrExpr) String() string {
	return fmt.Sprintf("%v - %v", x.Left, x.Right)
}

func (n SubtrExpr) Reducible() bool {
	return true
}

func (expr SubtrExpr) Reduce(env map[string]Expr) Expr {
	if expr.Left.Reducible() {
		return SubtrExpr{Left: expr.Left.Reduce(env), Right: expr.Right}
	} else if expr.Right.Reducible() {
		return SubtrExpr{Left: expr.Left, Right: expr.Right.Reduce(env)}
	} else {
		return Number{Val: expr.Left.(Number).Val - expr.Right.(Number).Val}
	}
}

type DivExpr struct {
	Left, Right Expr
}

func (x DivExpr) String() string {
	return fmt.Sprintf("%v / %v", x.Left, x.Right)
}

func (n DivExpr) Reducible() bool {
	return true
}

func (expr DivExpr) Reduce(env map[string]Expr) Expr {
	if expr.Left.Reducible() {
		return DivExpr{Left: expr.Left.Reduce(env), Right: expr.Right}
	} else if expr.Right.Reducible() {
		return DivExpr{Left: expr.Left, Right: expr.Right.Reduce(env)}
	} else {
		return Number{Val: expr.Left.(Number).Val / expr.Right.(Number).Val}
	}
}
