package simplex

import (
	"testing"
)

func Test_NestedMultExpr(t *testing.T) {
	three := Number{3}
	four := Number{4}
	five := Number{5}
	six := Number{6}
	a := AddExpr{
		MultExpr{three, four},
		MultExpr{five, six},
	}
	var hm Expr
	hm = a
	env := make(map[string]Expr)
	for hm.Reducible() {
		hm = hm.Reduce(env)
	}
	result := hm.(Number).Val
	want := 42
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_Machine(t *testing.T) {
	three := Number{3}
	four := Number{4}
	five := Number{5}
	six := Number{6}
	a := AddExpr{
		MultExpr{three, four},
		MultExpr{five, six},
	}
	env := make(map[string]Expr)
	m := Machine{a, env}
	m.Run(true)
	result := m.Expr.(Number).Val
	want := 42
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_LessThan(t *testing.T) {
	three := Number{3}
	four := Number{4}
	x := LessThanExpr{four, three}
	env := make(map[string]Expr)
	m := Machine{x, env}
	m.Run(true)
	result := m.Expr.(Boolean).Val
	want := false
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_AddWithLessThan(t *testing.T) {
	three := Number{3}
	four := Number{4}
	ten := Number{10}
	x := LessThanExpr{AddExpr{four, three}, ten}
	env := make(map[string]Expr)
	m := Machine{x, env}
	m.Run(true)
	result := m.Expr.(Boolean).Val
	want := true
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_WithEnv(t *testing.T) {
	x := AddExpr{Variable{Name: "x"}, Variable{Name: "y"}}
	env := make(map[string]Expr)
	env["x"] = Number{3}
	env["y"] = Number{4}
	m := Machine{x, env}
	m.Run(true)
	result := m.Expr.(Number).Val
	want := 7
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_Assignment(t *testing.T) {
	stmt := Assignment{Name: "x", X: AddExpr{Variable{Name: "x"}, Number{Val: 3}}}
	env := make(map[string]Expr)
	env["x"] = Number{3}
	var gen Stmt
	var newEnv map[string]Expr
	gen = stmt
	newEnv = env
	for gen.Reducible() {
		t.Log(gen)
		gen, newEnv = gen.Reduce(newEnv)
	}
	t.Log(gen)
	t.Logf("New env: %v", newEnv)
	result := newEnv["x"].(Number).Val
	want := 6
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()

	}
}

func Test_StmtMachine(t *testing.T) {
	stmt := Assignment{Name: "x", X: AddExpr{Variable{Name: "x"}, Number{Val: 3}}}
	env := make(map[string]Expr)
	env["x"] = Number{17}
	m := StmtMachine{stmt, env}
	m.Run(false)
	result := m.Env["x"].(Number).Val
	want := 20
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()

	}
}

func Test_ConditionalTrue(t *testing.T) {
	stmt := Conditional{
		Condition: Variable{
			Name: "x",
		},
		Consequence: Assignment{
			Name: "y",
			X:    Number{Val: 1},
		},
		Alternative: Assignment{
			Name: "y",
			X:    Number{Val: 2},
		},
	}
	env := make(map[string]Expr)
	env["x"] = Boolean{Val: true}
	m := StmtMachine{stmt, env}
	m.Run(false)
	result := m.Env["y"].(Number).Val
	want := 1
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_ConditionalFalse(t *testing.T) {
	stmt := Conditional{
		Condition: Variable{
			Name: "x",
		},
		Consequence: Assignment{
			Name: "y",
			X:    Number{Val: 1},
		},
		Alternative: Assignment{
			Name: "y",
			X:    Number{Val: 2},
		},
	}
	env := make(map[string]Expr)
	env["x"] = Boolean{Val: false}
	m := StmtMachine{stmt, env}
	m.Run(false)
	result := m.Env["y"].(Number).Val
	want := 2
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_Sequence(t *testing.T) {
	stmt0 := Assignment{Name: "x", X: AddExpr{Number{Val: 1}, Number{Val: 1}}}
	stmt1 := Assignment{Name: "y", X: AddExpr{Variable{Name: "x"}, Number{Val: 3}}}
	env := make(map[string]Expr)
	ss := Sequence{First: stmt0, Second: stmt1}
	m := StmtMachine{ss, env}
	m.Run(false)
	result := m.Env["y"].(Number).Val
	want := 5
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}

func Test_While(t *testing.T) {
	s := While{
		Condition: LessThanExpr{
			Left:  Variable{Name: "x"},
			Right: Number{5},
		},
		Body: Assignment{
			Name: "x",
			X:    AddExpr{Left: Variable{Name: "x"}, Right: Number{1}},
		},
	}
	env := make(map[string]Expr)
	env["x"] = Number{0}
	m := StmtMachine{s, env}
	m.Run(false)
	result := m.Env["x"].(Number).Val
	want := 5
	if result != want {
		t.Logf("wanted: %v, got: %v", want, result)
		t.Fail()
	}
}
