package main

import (
	"fmt"
	"github.com/kitasuna/uncomp/simplex"
)

func main() {
	m := simplex.Number{Val: 3}
	n := simplex.Number{Val: 4}

	addExpr := simplex.AddExpr{Left: m, Right: n}

	env := make(map[string]simplex.Expr)
	result := addExpr.Reduce(env)

	fmt.Printf("Result: %v\n", result)
}
