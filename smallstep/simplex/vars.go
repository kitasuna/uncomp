package simplex

type Variable struct {
	Name string
}

func (v Variable) String() string {
	return v.Name
}

func (v Variable) Reducible() bool {
	return true
}

func (v Variable) Reduce(environment map[string]Expr) Expr {
	return environment[v.Name]
}
