package parser

import (
	"text/scanner"
)

var keywords = map[string]int{
	"func":  FUNC,
	"class": CLASS,
	"if":    IF,
	"true":  TRUE,
	"false": FALSE,
	"nil":   NIL,
}

type Lexer struct {
	scanner.Scanner
	result Expr
	line   int
	col    int
}

func NewLexer() (l *Lexer) {
	l = &Lexer{}
	l.Scanner.Mode = scanner.GoTokens
	return
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())

	if token == scanner.Ident {
		if keyword, ok := keywords[l.TokenText()]; ok {
			token = keyword
		} else {
			token = IDENTIFIER
		}
	} else if token == scanner.Int || token == scanner.Float {
		token = NUMBER
	} else if token == scanner.String {
		token = STRING
	}

	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}
