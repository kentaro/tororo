%{
package parser

import (
    "io"
)

func Parse(io io.Reader) Expr {
    l := NewLexer()
    l.Init(io)
    yyParse(l)
    return l.result
}

%}

%union{
    stmt  Stmt
    expr  Expr
    token Token
}

%type<stmt> program
%type<expr> expr

// keywords
%token<token> FUNC CLASS
%token<token> IF
%token<token> TRUE FALSE NIL
%token<token> IDENTIFIER

// types
%token<token> NUMBER
%token<token> STRING

%left  '+', '-'
%left  '*', '/'

%%

program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

expr
    : NUMBER
    {
        $$ = Number{Literal: $1.Literal}
    }
    | STRING
    {
        $$ = String{Literal: $1.Literal}
    }
    | TRUE
    {
        $$ = true
    }
    | FALSE
    {
        $$ = false
    }
    | NIL
    {
        $$ = nil
    }
    | expr '+' expr
    {
        $$ = BinOp{Left: $1, Operator: '+', Right: $3}
    }
    | expr '-' expr
    {
        $$ = BinOp{Left: $1, Operator: '-', Right: $3}
    }
    | expr '*' expr
    {
        $$ = BinOp{Left: $1, Operator: '*', Right: $3}
    }
    | expr '/' expr
    {
        $$ = BinOp{Left: $1, Operator: '/', Right: $3}
    }
    | IF expr '{' expr '}'
    {
        // TODO
        $$ = If{Expr: $4}
    }

%%
