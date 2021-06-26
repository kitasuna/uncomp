package simplex

type Expr interface {
	Reduce(env map[string]Expr) Expr
	Reducible() bool
}

type Stmt interface {
	Reduce(env map[string]Expr) (Stmt, map[string]Expr)
	Reducible() bool
	Equals(other Stmt) bool
}
