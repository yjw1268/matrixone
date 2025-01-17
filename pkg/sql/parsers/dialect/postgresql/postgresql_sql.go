// Code generated by goyacc -o postgresql_sql.go postgresql_sql.y. DO NOT EDIT.

//line postgresql_sql.y:16
package postgresql

import (
	__yyfmt__ "fmt"
	__yyunsafe__ "unsafe"
)

//line postgresql_sql.y:16

import (
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/tree"
)

const LEX_ERROR = 57346
const EMPTY = 57347
const UNION = 57348
const SELECT = 57349
const STREAM = 57350
const INSERT = 57351
const UPDATE = 57352
const DELETE = 57353
const FROM = 57354
const WHERE = 57355
const GROUP = 57356
const HAVING = 57357
const ORDER = 57358
const BY = 57359
const LIMIT = 57360
const OFFSET = 57361
const FOR = 57362
const LOWER_THAN_SET = 57363
const SET = 57364
const ALL = 57365
const DISTINCT = 57366
const DISTINCTROW = 57367
const AS = 57368
const EXISTS = 57369
const ASC = 57370
const DESC = 57371
const INTO = 57372
const DUPLICATE = 57373
const DEFAULT = 57374
const LOCK = 57375
const KEYS = 57376
const VALUES = 57377
const NEXT = 57378
const VALUE = 57379
const SHARE = 57380
const MODE = 57381
const SQL_NO_CACHE = 57382
const SQL_CACHE = 57383
const JOIN = 57384
const STRAIGHT_JOIN = 57385
const LEFT = 57386
const RIGHT = 57387
const INNER = 57388
const OUTER = 57389
const CROSS = 57390
const NATURAL = 57391
const USE = 57392
const FORCE = 57393
const ON = 57394
const USING = 57395
const SUBQUERY_AS_EXPR = 57396
const LOWER_THAN_STRING = 57397
const ID = 57398
const AT_ID = 57399
const AT_AT_ID = 57400
const STRING = 57401
const VALUE_ARG = 57402
const LIST_ARG = 57403
const COMMENT = 57404
const COMMENT_KEYWORD = 57405
const INTEGRAL = 57406
const HEX = 57407
const BIT_LITERAL = 57408
const FLOAT = 57409
const HEXNUM = 57410
const NULL = 57411
const TRUE = 57412
const FALSE = 57413
const LOWER_THAN_CHARSET = 57414
const CHARSET = 57415
const UNIQUE = 57416
const KEY = 57417
const OR = 57418
const PIPE_CONCAT = 57419
const XOR = 57420
const AND = 57421
const NOT = 57422
const BETWEEN = 57423
const CASE = 57424
const WHEN = 57425
const THEN = 57426
const ELSE = 57427
const END = 57428
const LE = 57429
const GE = 57430
const NE = 57431
const NULL_SAFE_EQUAL = 57432
const IS = 57433
const LIKE = 57434
const REGEXP = 57435
const IN = 57436
const ASSIGNMENT = 57437
const SHIFT_LEFT = 57438
const SHIFT_RIGHT = 57439
const DIV = 57440
const MOD = 57441
const UNARY = 57442
const COLLATE = 57443
const BINARY = 57444
const UNDERSCORE_BINARY = 57445
const INTERVAL = 57446

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"EMPTY",
	"UNION",
	"SELECT",
	"STREAM",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"OFFSET",
	"FOR",
	"LOWER_THAN_SET",
	"SET",
	"ALL",
	"DISTINCT",
	"DISTINCTROW",
	"AS",
	"EXISTS",
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"DEFAULT",
	"LOCK",
	"KEYS",
	"VALUES",
	"NEXT",
	"VALUE",
	"SHARE",
	"MODE",
	"SQL_NO_CACHE",
	"SQL_CACHE",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"USING",
	"SUBQUERY_AS_EXPR",
	"'('",
	"','",
	"')'",
	"LOWER_THAN_STRING",
	"ID",
	"AT_ID",
	"AT_AT_ID",
	"STRING",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"COMMENT_KEYWORD",
	"INTEGRAL",
	"HEX",
	"BIT_LITERAL",
	"FLOAT",
	"HEXNUM",
	"NULL",
	"TRUE",
	"FALSE",
	"LOWER_THAN_CHARSET",
	"CHARSET",
	"UNIQUE",
	"KEY",
	"OR",
	"PIPE_CONCAT",
	"XOR",
	"AND",
	"NOT",
	"'!'",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"ASSIGNMENT",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"DIV",
	"'%'",
	"MOD",
	"'^'",
	"'~'",
	"UNARY",
	"COLLATE",
	"BINARY",
	"UNDERSCORE_BINARY",
	"INTERVAL",
	"'.'",
	"';'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line postgresql_sql.y:110

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 11

var yyAct = [...]int{
	6, 7, 5, 3, 1, 4, 2, 0, 0, 0,
	8,
}

var yyPact = [...]int{
	-48, -1000, -122, -1000, -1000, -58, -48, -1000, -1000,
}

var yyPgo = [...]int{
	0, 3, 6, 5, 4,
}

//line postgresql_sql.y:110
type yySymType struct {
	union interface{}
	id    int
	str   string
	item  interface{}
	yys   int
}

func (st *yySymType) statementUnion() tree.Statement {
	v, _ := st.union.(tree.Statement)
	return v
}

func (st *yySymType) statementsUnion() []tree.Statement {
	v, _ := st.union.([]tree.Statement)
	return v
}

var yyR1 = [...]int{
	0, 4, 2, 2, 1, 3, 3,
}

var yyR2 = [...]int{
	0, 1, 1, 3, 1, 2, 1,
}

var yyChk = [...]int{
	-1000, -4, -2, -1, -3, 50, 122, 59, -1,
}

var yyDef = [...]int{
	0, -2, 1, 2, 4, 6, 0, 5, 3,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 84, 3, 3, 3, 112, 104, 3,
	55, 57, 109, 107, 56, 108, 121, 110, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 122,
	92, 91, 93, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 114, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 103, 3, 115,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 82, 83, 85,
	86, 87, 88, 89, 90, 94, 95, 96, 97, 98,
	99, 100, 101, 102, 105, 106, 111, 113, 116, 117,
	118, 119, 120,
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

func yyIaddr(v interface{}) __yyunsafe__.Pointer {
	type h struct {
		t __yyunsafe__.Pointer
		p __yyunsafe__.Pointer
	}
	return (*h)(__yyunsafe__.Pointer(&v)).p
}

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

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line postgresql_sql.y:86
		{
			if yyDollar[1].statementUnion() != nil {
				yylex.(*Lexer).AppendStmt(yyDollar[1].statementUnion())
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line postgresql_sql.y:92
		{
			if yyDollar[3].statementUnion() != nil {
				yylex.(*Lexer).AppendStmt(yyDollar[3].statementUnion())
			}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		var yyLOCAL tree.Statement
//line postgresql_sql.y:103
		{
			yyLOCAL = &tree.Use{Name: yyDollar[2].str}
		}
		yyVAL.union = yyLOCAL
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		var yyLOCAL tree.Statement
//line postgresql_sql.y:107
		{
			yyLOCAL = &tree.Use{}
		}
		yyVAL.union = yyLOCAL
	}
	goto yystack /* stack new state and value */
}
