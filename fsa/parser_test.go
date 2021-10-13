package main

import (
	"github.com/alecthomas/participle"
	"testing"
)


func Test_ParserHandlesSingleRune(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("a", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesMultipleRunes(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("ab", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesConditional(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("a|b", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesRepetition(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("ab*", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesBrackets(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("(ab)", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesBracketsWithChooseAndRepeat(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("(a|b)*", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesNestedRepeat(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("a|b*", ast)
	if err != nil {
		t.Error(err)
	}
}


func Test_ParserHandlesRepetitionWitBrackets(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("(a)*", ast)
	if err != nil {
		t.Error(err)
	}
}

func Test_ParserHandlesChooseWithEmpty(t *testing.T) {
	parser, err := participle.Build(&PRegex{})
	if err != nil {
		t.Error(err)
	}

	ast := &PRegex{}
	err = parser.ParseString("nil|b", ast)
	if err != nil {
		t.Error(err)
	}
}
