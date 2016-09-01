package parser

type Stmt interface{}
type Expr interface{}

type Token struct {
	Token   int
	Literal string
}

type Number struct {
	Literal string
}

type String struct {
	Literal string
}

type BinOp struct {
	Left     Expr
	Operator rune
	Right    Expr
}

// TODO: must be moved to statements
type If struct {
	Expr Expr
}
