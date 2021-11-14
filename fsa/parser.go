package main

/*
import (
 "github.com/alecthomas/participle"
)
*/

/*
type PartcRegex struct {
	Elems []*Elem `@@*`
}
*/

type PRegex struct {
	Lhs  *Value  `@@`
	Tail []*Oper `@@*`
}

type Value struct {
	Id      *string `@Ident`
	SubExpr *PRegex `| "(" @@ ")"`
	// Not ideal, but leaving this as a placeholder to represent "Empty"
	Nothing bool `| @"nil"`
}

type Oper struct {
	Choosey string `"|" @Ident`
	// Rhs *Value "@@"
	Repeaty string `| "*"`
	// SubElem *PRegex `| "(" @@ ")"`
}

type ChooseEx struct {
	First    *PRegex `@@`
	Operator string  `"|"`
	Second   *PRegex `@@`
}

type RepeatEx struct {
	First    *PRegex `@@`
	Operator string  `"*"`
}

/*
type AddExpr struct {
  Left *Term `@@`
  Operator string ` "+" `
  Right *Term `@@`
}
*/
