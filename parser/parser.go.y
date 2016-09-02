%{
package parser

import (
    "io"
)

func Parse(io io.Reader) []Stmt {
    l := NewLexer()
    l.Init(io)
    yyParse(l)
    return l.result
}

%}

%union{
    stmts []Stmt
    stmt  Stmt
    exprs []Expr
    expr  Expr
    token Token
}

%type<stmts> program
%type<stmts> stmts
%type<stmt>  stmt
%type<exprs> exprs
%type<expr>  expr

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
    :
    {
        $$ = nil
    }
    | stmts
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

stmts
    : stmt
    {
        $$ = []Stmt{$1}
    }
    | stmts stmt
    {
        $$ = append($1, $2)
    }

stmt
    : exprs
    {
        $$ = $1
    }
    | FUNC expr '(' /* TODO: VAR */ ')' '{' stmts '}'
    {
        $$ = FuncStmt{
            Name:  $2.(IdentifierExpr).Literal,
            Args:  []string{},
            Stmts: $6,
        }
    }
    | CLASS expr '{' stmts '}'
    {
        $$ = ClassStmt{
            Name:  $2.(IdentifierExpr).Literal,
            Stmts: $4,
        }
    }
    | IF expr '{' stmts '}'
    {
        $$ = IfStmt{
            Cond: $2,
            Then: $4,
        }
    }

exprs
    : expr
    {
        $$ = []Expr{$1}
    }
    | exprs expr
    {
        $$ = append($1, $2)
    }

expr
    : NUMBER
    {
        $$ = NumberExpr{Literal: $1.Literal}
    }
    | STRING
    {
        $$ = StringExpr{Literal: $1.Literal}
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
    | IDENTIFIER
    {
        $$ = IdentifierExpr{Literal: $1.Literal}
    }
    | expr '+' expr
    {
        $$ = BinOpExpr{Left: $1, Operator: '+', Right: $3}
    }
    | expr '-' expr
    {
        $$ = BinOpExpr{Left: $1, Operator: '-', Right: $3}
    }
    | expr '*' expr
    {
        $$ = BinOpExpr{Left: $1, Operator: '*', Right: $3}
    }
    | expr '/' expr
    {
        $$ = BinOpExpr{Left: $1, Operator: '/', Right: $3}
    }

%%
