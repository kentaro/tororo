package parser

import (
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	result Expr
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = NUMBER
	}
	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}
