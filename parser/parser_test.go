package parser

import (
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
			NumberExpr{
				Literal: "9999",
			},
		},
	},
	{
		"9.999",
		[]Stmt{
			NumberExpr{
				Literal: "9.999",
			},
		},
	},
	{
		"1 + 1",
		[]Stmt{
			BinOpExpr{
				Left:     NumberExpr{Literal: "1"},
				Operator: '+',
				Right:    NumberExpr{Literal: "1"},
			},
		},
	},
	{
		`"foo \"bar\" baz"`,
		[]Stmt{
			StringExpr{
				Literal: `"foo \"bar\" baz"`,
			},
		},
	},
	// TODO: test
	// {
	// 	`
	// 	class Foo {
	// 		"got true"
	// 	}`,
	// 	[]Stmt{
	// 		ClassStmt{
	// 			Name: "Foo",
	// 			Stmts: []Stmt{[]Expr{StringExpr{Literal: "got true"}}},
	// 		},
	// 	},
	// },
}

func TestParse(t *testing.T) {
	for i, test := range tests {
		stmts := Parse(strings.NewReader(test.src))

		for j, stmt := range stmts {
			switch stmt.(type) {
			case []Expr:
				for _, expr := range stmt.([]Expr) {
					if expr != test.ast[j] {
						t.Errorf("case %d:\n\n%s\n\nactual:\t\t%#v\nexpected:\t%#v", i, test.src, expr, test.ast[j])
					}
				}
			default:
				if stmt != test.ast[j] {
					t.Errorf("case %d:\n\n%s\n\nactual:\t\t%#v\nexpected:\t%#v", i, test.src, stmt, test.ast[j])
				}
			}
		}
	}
}
