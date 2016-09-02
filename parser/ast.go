package parser

type Stmt interface{}

type FuncStmt struct {
	Name  string
	Args  []string
	Stmts []Stmt
}

type ClassStmt struct {
	Name  string
	Stmts []Stmt
}

type IfStmt struct {
	Cond Expr
	Then []Stmt
}

type Expr interface{}

type NumberExpr struct {
	Literal string
}

type StringExpr struct {
	Literal string
}

type IdentifierExpr struct {
	Literal string
}

type BinOpExpr struct {
	Left     Expr
	Operator rune
	Right    Expr
}

type Token struct {
	Token   int
	Literal string
}
