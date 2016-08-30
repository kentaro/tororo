package parser

type Expr interface{}

type Token struct {
	Token   int
	Literal string
}

type Number struct {
	Literal string
}

type BinOp struct {
	Left     Expr
	Operator rune
	Right    Expr
}
