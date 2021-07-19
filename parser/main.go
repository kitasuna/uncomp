package main

import (
 "github.com/alecthomas/participle"
 "bytes"
 "fmt"
 "log"
 "os"
)

type SimpleProg struct {
  Statements []*Statement `@@*`
}

func (sp SimpleProg) ToJS() string {
  var buff bytes.Buffer
  for _, stmt := range sp.Statements {
    buff.WriteString(stmt.ToJS())
  }
  return buff.String()
}

type Statement struct {
  Assignment *Assignment `@@`
  While *While `| @@`
}

func (s Statement) ToJS() string {
  if s.Assignment != nil {
    return s.Assignment.ToJS()
  } else if s.While != nil {
    return s.While.ToJS()
  }

  return "<<STATEMENT>>"
}

type Assignment struct {
  Name string `@Ident "="`
  X *Expr `@@`
}
func (a Assignment) ToJS() string {
	return fmt.Sprintf("(env) => Object.assign(env, {'%v': (" + a.X.ToJS() + ")(env)})", a.Name)
}

type Expr struct {
  AddExpr *AddExpr ` @@`
  MultExpr *MultExpr `| @@`
  LTExpr *LTExpr `| @@`
  Number *int `| @Int`
}

func (x *Expr) ToJS() string {
  if x.AddExpr != nil {
    return x.AddExpr.ToJS()
  } else if x.MultExpr != nil {
    return x.MultExpr.ToJS()
  } else if x.LTExpr != nil {
    return x.LTExpr.ToJS()
  } else if x.Number != nil {
    return fmt.Sprint(*x.Number)
  }

  return "<<EXPR>>"
}

type Variable struct {
  Name string `@Ident`
}

func (v Variable) ToJS() string {
	return fmt.Sprintf("(env) => env['%v']", v.Name)
}

type Term struct {
  Number *int `@Int`
  Var *Variable `|@@`
}

func (t Term) ToJS() string {
  if t.Number != nil {
    return fmt.Sprintf("(env) => %v", *(t.Number))
  } else if t.Var != nil {
    return t.Var.ToJS()
  }

  return "<<TERM>>"
}

type AddExpr struct {
  Left *Term `@@`
  Operator string ` "+" `
  Right *Term `@@`
}

func (x AddExpr) ToJS() string {
	return "(env) => (" + x.Left.ToJS() + ")(env) + (" + x.Right.ToJS() + ")(env)"
}


type MultExpr struct {
  Left *Term `@@`
  Operator string ` "*" `
  Right *Term `@@`
}

func (x MultExpr) ToJS() string {
	return "(env) => (" + x.Left.ToJS() + ")(env) * (" + x.Right.ToJS() + ")(env)"
}


type LTExpr struct {
  Left *Term `@@`
  Operator string ` "<" `
  Right *Term `@@`
}

func (x LTExpr) ToJS() string {
	return "(env) =>  (" + x.Left.ToJS() + ")(env) < (" + x.Right.ToJS() + ")(env)"
}

type While struct {
  Condition *Expr `"while" "("@@")"`
  Body *Statement `"{" @@ "}"`
}
func (w While) ToJS() string {
  return "(env) => { while (("+ w.Condition.ToJS() +")(env)) { env = (" + w.Body.ToJS() + ")(env); } return env;}"
}


func main() {
  parser, err := participle.Build(&SimpleProg{})
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  ast := &SimpleProg{}
  err = parser.ParseString("while (x < 3) {\nx = x + 1\n}", ast)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  fmt.Println(ast.ToJS())
}
