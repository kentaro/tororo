%{
package parser

import (
    "io"
)

func Parse(io io.Reader) Expr {
    l := &Lexer{}
    l.Init(io)
    yyParse(l)
    return l.result
}

%}

%union{
    token Token
    expr  Expr
}

%type<expr>   program
%type<expr>   expr
%token<token> NUMBER

%left '+', '-'
%left '*', '/'

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

%%
