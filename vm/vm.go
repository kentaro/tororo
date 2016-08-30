package vm

import (
	"strconv"
	"github.com/kentaro/tororo/parser"
)

var insnsMap = map[string](func(int, int) int){
	"+": Add,
	"-": Subtract,
	"*": Multiply,
	"/": Divide,
}

func Run(e parser.Expr) int {
	return eval(e)
}

func eval(e parser.Expr) (result int) {
	switch e.(type) {
		case parser.BinOp:
			left := eval(e.(parser.BinOp).Left)
			right := eval(e.(parser.BinOp).Right)
			result = insnsMap[string(e.(parser.BinOp).Operator)](left, right)
		case parser.Number:
			num, _ := strconv.Atoi(e.(parser.Number).Literal)
			result = num
	}

	return
}
