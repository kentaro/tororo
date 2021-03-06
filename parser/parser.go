//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"io"
)

func Parse(io io.Reader) []Stmt {
	l := NewLexer()
	l.Init(io)
	yyParse(l)
	return l.result
}

//line parser.go.y:17
type yySymType struct {
	yys   int
	stmts []Stmt
	stmt  Stmt
	exprs []Expr
	expr  Expr
	token Token
}

const FUNC = 57346
const CLASS = 57347
const IF = 57348
const TRUE = 57349
const FALSE = 57350
const NIL = 57351
const IDENTIFIER = 57352
const NUMBER = 57353
const STRING = 57354

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"FUNC",
	"CLASS",
	"IF",
	"TRUE",
	"FALSE",
	"NIL",
	"IDENTIFIER",
	"NUMBER",
	"STRING",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"'{'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:147

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 21
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 101

var yyAct = [...]int{

	3, 2, 8, 15, 34, 31, 4, 16, 17, 18,
	19, 22, 23, 5, 6, 7, 11, 12, 13, 14,
	9, 10, 1, 27, 28, 29, 30, 32, 33, 38,
	0, 0, 0, 15, 15, 0, 37, 0, 15, 5,
	6, 7, 11, 12, 13, 14, 9, 10, 0, 20,
	21, 22, 23, 24, 0, 36, 5, 6, 7, 11,
	12, 13, 14, 9, 10, 0, 0, 20, 21, 22,
	23, 0, 35, 26, 20, 21, 22, 23, 0, 0,
	25, 5, 6, 7, 11, 12, 13, 14, 9, 10,
	0, 20, 21, 22, 23, 11, 12, 13, 14, 9,
	10,
}
var yyPact = [...]int{

	77, -1000, 77, -1000, 88, 88, 88, 88, 78, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 78, 36, 61, 54,
	88, 88, 88, 88, -13, 77, 77, -4, -4, -1000,
	-1000, -15, 52, 35, 77, -1000, -1000, 9, -1000,
}
var yyPgo = [...]int{

	0, 22, 1, 0, 6, 2,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 4,
	4, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5,
}
var yyR2 = [...]int{

	0, 0, 1, 1, 2, 1, 7, 5, 5, 1,
	2, 1, 1, 1, 1, 1, 1, 3, 3, 3,
	3,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 4, 5, 6, -5, 11,
	12, 7, 8, 9, 10, -3, -5, -5, -5, -5,
	13, 14, 15, 16, 17, 19, 19, -5, -5, -5,
	-5, 18, -2, -2, 19, 20, 20, -2, 20,
}
var yyDef = [...]int{

	1, -2, 2, 3, 5, 0, 0, 0, 9, 11,
	12, 13, 14, 15, 16, 4, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 17, 18, 19,
	20, 0, 0, 0, 0, 7, 8, 0, 6,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	17, 18, 15, 13, 3, 14, 3, 16, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 3, 20,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:48
		{
			yyVAL.stmts = nil
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:52
		{
			yyVAL.stmts = yyDollar[1].stmts
			yylex.(*Lexer).result = yyVAL.stmts
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:59
		{
			yyVAL.stmts = []Stmt{yyDollar[1].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:63
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.stmt = yyDollar[1].exprs
		}
	case 6:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.stmt = FuncStmt{
				Name:  yyDollar[2].expr.(IdentifierExpr).Literal,
				Args:  []string{},
				Stmts: yyDollar[6].stmts,
			}
		}
	case 7:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.stmt = ClassStmt{
				Name:  yyDollar[2].expr.(IdentifierExpr).Literal,
				Stmts: yyDollar[4].stmts,
			}
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:88
		{
			yyVAL.stmt = IfStmt{
				Cond: yyDollar[2].expr,
				Then: yyDollar[4].stmts,
			}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:101
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.expr = NumberExpr{Literal: yyDollar[1].token.Literal}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.expr = StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:115
		{
			yyVAL.expr = true
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.expr = false
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.expr = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.expr = IdentifierExpr{Literal: yyDollar[1].token.Literal}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.expr = BinOpExpr{Left: yyDollar[1].expr, Operator: '+', Right: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.expr = BinOpExpr{Left: yyDollar[1].expr, Operator: '-', Right: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.expr = BinOpExpr{Left: yyDollar[1].expr, Operator: '*', Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.expr = BinOpExpr{Left: yyDollar[1].expr, Operator: '/', Right: yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
