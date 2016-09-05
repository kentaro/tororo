package parser

import (
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	src string
	ast []Stmt
}{
	{
		"9999",
		[]Stmt{
			[]Expr{NumberExpr{
				Literal: "9999",
			}},
		},
	},
	{
		"9.999",
		[]Stmt{
			[]Expr{NumberExpr{
				Literal: "9.999",
			}},
		},
	},
	{
		"foo",
		[]Stmt{
			[]Expr{IdentifierExpr{
				Literal: "foo",
			}},
		},
	},
	{
		"1 + 1",
		[]Stmt{
			[]Expr{BinOpExpr{
				Left:     NumberExpr{Literal: "1"},
				Operator: '+',
				Right:    NumberExpr{Literal: "1"},
			}},
		},
	},
	{
		"1 - 1",
		[]Stmt{
			[]Expr{BinOpExpr{
				Left:     NumberExpr{Literal: "1"},
				Operator: '-',
				Right:    NumberExpr{Literal: "1"},
			}},
		},
	},
	{
		"1 * 1",
		[]Stmt{
			[]Expr{BinOpExpr{
				Left:     NumberExpr{Literal: "1"},
				Operator: '*',
				Right:    NumberExpr{Literal: "1"},
			}},
		},
	},
	{
		"1 / 1",
		[]Stmt{
			[]Expr{BinOpExpr{
				Left:     NumberExpr{Literal: "1"},
				Operator: '/',
				Right:    NumberExpr{Literal: "1"},
			}},
		},
	},
	{
		`"foo \"bar\" baz"`,
		[]Stmt{
			[]Expr{StringExpr{
				Literal: `"foo \"bar\" baz"`,
			}},
		},
	},
	{
		`
		func foo() {
			"bar"
			"baz"
		}`,
		[]Stmt{
			FuncStmt{
				Name: "foo",
				Args: []string{},
				Stmts: []Stmt{[]Expr{
					StringExpr{Literal: "\"bar\""},
					StringExpr{Literal: "\"baz\""},
				}},
			},
		},
	},
	{
		`
		class Foo {
			"bar"
			"baz"
		}`,
		[]Stmt{
			ClassStmt{
				Name: "Foo",
				Stmts: []Stmt{[]Expr{
					StringExpr{Literal: "\"bar\""},
					StringExpr{Literal: "\"baz\""},
				}},
			},
		},
	},
	{
		`
		if true {
			"bar"
			"baz"
		}`,
		[]Stmt{
			IfStmt{
				Cond: true,
				Then: []Stmt{[]Expr{
					StringExpr{Literal: "\"bar\""},
					StringExpr{Literal: "\"baz\""},
				}},
			},
		},
	},
}

func TestParse(t *testing.T) {
	for i, test := range tests {
		stmts := Parse(strings.NewReader(test.src))
		if !reflect.DeepEqual(stmts, test.ast) {
			t.Errorf("case %d:\n\n%s\n\nactual:\t\t%#v\nexpected:\t%#v", i, test.src, stmts, test.ast)
		}
	}
}
