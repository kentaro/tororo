package parser

import (
	"strings"
	"testing"
)

var tests = []struct {
	src string
	ast interface{}
}{
	{
		"9999",
		Number{
			Literal: "9999",
		},
	},
	{
		"9.999",
		Number{
			Literal: "9.999",
		},
	},
	{
		"1 + 1",
		BinOp{
			Left:     Number{Literal: "1"},
			Operator: '+',
			Right:    Number{Literal: "1"},
		},
	},
	{
		`"foo \"bar\" baz"`,
		String{
			Literal: `"foo \"bar\" baz"`,
		},
	},
	{
		`
		if true {
			"got true"
		}`,
		If{Expr: String{Literal: "\"got true\""}},
	},
}

func TestParse(t *testing.T) {
	for i, test := range tests {
		res := Parse(strings.NewReader(test.src))
		if res != test.ast {
			t.Errorf("case %d:\n\n%s\n\nactual:\t\t%#v\nexpected:\t%#v", i, test.src, res, test.ast)
		}
	}
}
