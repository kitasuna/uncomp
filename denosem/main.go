package main

import (
	"fmt"
  "io/ioutil"
)

type Env map[string]interface{}

type DenoExpr interface {
	ToJS() string
}

type Number struct {
	Val int
}

func (n Number) ToJS() string {
	return fmt.Sprintf("(env) => %v", n.Val)
}

type Boolean struct {
	Val bool
}

func (b Boolean) ToJS() string {
	return fmt.Sprintf("(env) => %v", b.Val)
}

type Variable struct {
	Name string
}

func (v Variable) ToJS() string {
	return fmt.Sprintf("(env) => env['%v']", v.Name)
}

type AddExpr struct {
	Left, Right DenoExpr
}

func (x AddExpr) ToJS() string {
	return "(env) => (" + x.Left.ToJS() + ")(env) + (" + x.Right.ToJS() + ")(env)"
}

type MultExpr struct {
	Left, Right DenoExpr
}

func (x MultExpr) ToJS() string {
	return "(env) => (" + x.Left.ToJS() + ")(env) * (" + x.Right.ToJS() + ")(env)"
}

type LessThanExpr struct {
	Left, Right DenoExpr
}

func (x LessThanExpr) ToJS() string {
	return "(env) =>  (" + x.Left.ToJS() + ")(env) < (" + x.Right.ToJS() + ")(env)"
}

type Statement interface {
	IsStatement() bool
	ToJS() string
}
type Assignment struct {
	Name string
	X DenoExpr
}

func (a Assignment) ToJS() string {
	return fmt.Sprintf("(env) => Object.assign(env, {'%v': (" + a.X.ToJS() + ")(env)})", a.Name)
}
func (a Assignment) IsStatement() bool {
	return true
}

type Noop struct {}

func (n Noop) ToJS() string {
	return "(env) => env"
}

func (n Noop) IsStatement() bool {
	return true
}

type Conditional struct {
	Condition DenoExpr
	Consequence Statement
	Alternative Statement
}

func (c Conditional) ToJS() string {
	return "(env) => { if ((" + c.Condition.ToJS() + ")(env)) { return (" + c.Consequence.ToJS() + ")(env) } else { return (" + c.Alternative.ToJS() + ")(env) } }"
}

func (c Conditional) IsStatement() bool {
	return true
}

type Sequence struct {
  First Statement
  Second Statement
}

func (s Sequence) ToJS() string {
  return "(env) => { return (" + s.Second.ToJS() + ")((" + s.First.ToJS() + ")(env)) }"
}

func (s Sequence) IsStatement() bool {
  return true
}

type While struct {
  Condition DenoExpr
  Body Statement
}

func (w While) ToJS() string {
  return "(env) => { while (("+ w.Condition.ToJS() +")(env)) { env = (" + w.Body.ToJS() + ")(env); } return env;}"
}

func (w While) IsStatement() bool {
  return true
}
func main() {
	// In Node REPL:
	// g = Function("return " + fs.readFileSync('index.js').toString())
	addExpr := AddExpr{
		Left: Variable { Name: "x" },
		Right: Number { Val: 14 },
	}
  ioutil.WriteFile("./output/addexpr.js", []byte(addExpr.ToJS()), 0666)

	multExpr := MultExpr{
		Left: Variable { Name: "x" },
		Right: Number { Val: 14 },
	}
  ioutil.WriteFile("./output/multexpr.js", []byte(multExpr.ToJS()), 0666)

	lessThanExpr := LessThanExpr{
		Left: Variable { Name: "x" },
		Right: Number { Val: 14 },
	}
  ioutil.WriteFile("./output/lessthanexpr.js", []byte(lessThanExpr.ToJS()), 0666)

	assign := Assignment{
		Name: "z",
		X: AddExpr {
			Left: Number { Val: 32 },
			Right: Number { Val: 10 },
		},
	}
  ioutil.WriteFile("./output/assign.js", []byte(assign.ToJS()), 0666)

	cond := Conditional{
		Condition: LessThanExpr {
			Left: Variable {
				Name: "x",
			},
			Right: Number { Val: 42 },
		},
		Consequence: Assignment {
			Name: "lessThanFortyTwo",
			X: Boolean { Val: true },
		},
		Alternative: Assignment {
			Name: "moreThanFortyTwo",
			X: Boolean { Val: true },
		},
	}
  ioutil.WriteFile("./output/cond.js", []byte(cond.ToJS()), 0666)

  seq := Sequence{
    First: Assignment{
      Name: "z",
      X: AddExpr {
        Left: Number { Val: 3 },
        Right: Number { Val: 3 },
      },
    },
    Second: Conditional{
		Condition: LessThanExpr {
			Left: Variable {
				Name: "z",
			},
			Right: Number { Val: 42 },
		},
		Consequence: Assignment {
			Name: "lessThanFortyTwo",
			X: Boolean { Val: true },
		},
		Alternative: Assignment {
			Name: "moreThanFortyTwo",
			X: Boolean { Val: true },
		},
    },
  }
  ioutil.WriteFile("./output/seq.js", []byte(seq.ToJS()), 0666)

  while := While {
		Condition: LessThanExpr {
			Left: Variable {
				Name: "x",
			},
			Right: Number { Val: 42 },
		},
    Body: Sequence {
      First: Assignment {
        Name: "x",
        X: AddExpr { Left: Variable { Name: "x" } , Right: Number { Val: 1 } },
      },
      Second: Assignment {
        Name: "z",
        X: MultExpr { Left: Variable {Name: "x"}, Right: Number { Val: 2 } },
      },
    },
  }
  ioutil.WriteFile("./output/while.js", []byte(while.ToJS()), 0666)
}
