package main

import (
	"fmt"
)

type Env map[string]BigExpr

type BigExpr interface {
	Eval(env Env) BigExpr
}

type Number struct {
	Val int
}

func (n Number) Eval(env Env) BigExpr {
	return n
}

type Boolean struct {
	Val bool
}

func (b Boolean) Eval(env Env) BigExpr {
	return b
}

type Variable struct {
	Name string
}

func (v Variable) Eval(env Env) BigExpr {
	return env[v.Name]
}

type AddExpr struct {
	Left, Right BigExpr
}

func (x AddExpr) Eval(env Env) BigExpr {
	return Number {
		Val: x.Left.Eval(env).(Number).Val + x.Right.Eval(env).(Number).Val,
	}
}

type LessThanExpr struct {
	Left, Right BigExpr
}

func (x LessThanExpr) Eval(env Env) BigExpr {
	return Boolean {
		x.Left.Eval(env).(Number).Val < x.Right.Eval(env).(Number).Val,
	}
}

type Statement interface {
	Eval(env Env) Env
}

type Assignment struct {
	Name string
	X BigExpr
}

func (a Assignment) Eval(env Env) Env {
	env[a.Name] = a.X.Eval(env)
	return env
}

type Noop struct {}

func (n Noop) Eval(env Env) Env {
	return env
}

type Conditional struct {
	Condition BigExpr
	Consequence Statement
	Alternative Statement
}

func (c Conditional) Eval(env Env) Env {
	if c.Condition.Eval(env).(Boolean).Val == true {
		return c.Consequence.Eval(env)
	} else {
		return c.Alternative.Eval(env)
	}
}

type Seq struct {
	First Statement
	Second Statement
}

func (s Seq) Eval(env Env) Env {
	return s.Second.Eval(s.First.Eval(env))
}

type While struct {
	Condition BigExpr
	Body Statement
}

func (w While) Eval(env Env) Env {
	if w.Condition.Eval(env).(Boolean).Val == true {
		newEnv := w.Body.Eval(env)
		return w.Eval(newEnv)
	}

	return env
}

func main() {
	env := make(map[string]BigExpr)
	addX := AddExpr {
		Left: Number { 33 },
		Right: Number { 36 },
	}
	fmt.Printf("Add Result: %v\n", addX.Eval(env))
	ltX := LessThanExpr {
		Left: Number {100},
		Right: Number {101},
	}
	fmt.Printf("LT Result: %v\n", ltX.Eval(env))

	assX := Assignment {
		Name: "x"	,
		X: AddExpr {
			Left: Number { 3 },
			Right: Number {66},
		},
	}
	fmt.Printf("Ass result: %v\n", assX.Eval(env))

	yIsX := Assignment {
		Name: "y",
		X: Variable {
			Name: "x",
		},
	}

	seqX := Seq {
		First: assX,
		Second: yIsX,
	}
	fmt.Printf("Seq result: %v\n", seqX.Eval(env))

	whileX := While {
		Condition: LessThanExpr {
			Left: Variable { Name: "x" },
			Right: Number { Val: 15 },
		},
		Body: Assignment {
			Name: "x",
			X: AddExpr {
				Left: Variable { Name: "x" },
				Right: Number { Val: 1 },
			},
		},
	}
	fmt.Printf("While result: %v\n", whileX.Eval(Env{ "x": Number { Val: 0 } }))
}

