package token

import (
	"fmt"
	"strconv"
)

// 记号类型
type TokenType int

// ch3中 µGo程序用到的记号类型
const (
	EOF TokenType = iota // = 0
	ERROR
	COMMENT

	IDENT
	NUMBER

	PACKAGE
	IMPORT
	VAR
	FUNC
	RETURN
	IF
	ELSE
	FOR
	BREAK
	CONTINUE
	DEFER
	GOTO

	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	EQL // ==
	NEQ // !=
	LSS // <
	LEQ // <=
	GTR // >
	GEQ // >=

	ASSIGN // =
	DEFINE // :=

	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	COMMA     // ,
	SEMICOLON // ;
)

// 记号值
type Token struct {
	Pos     Pos       // 记号所在的位置(从1开始)
	Type    TokenType // 记号的类型
	Literal string    // 程序中原始的字符串
}

var tokens = [...]string{
	EOF:     "EOF",
	ERROR:   "ERROR",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	NUMBER: "NUMBER",

	PACKAGE:  "package",
	IMPORT:   "import",
	VAR:      "var",
	FUNC:     "func",
	RETURN:   "return",
	IF:       "if",
	ELSE:     "else",
	FOR:      "for",
	BREAK:    "break",
	CONTINUE: "continue",
	DEFER:    "defer",
	GOTO:     "goto",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	MOD: "%",

	EQL: "==",
	NEQ: "!=",
	LSS: "<",
	LEQ: "<=",
	GTR: ">",
	GEQ: ">=",

	ASSIGN: "=",
	DEFINE: ":=",

	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",

	COMMA:     ",",
	SEMICOLON: ";",
}

func (tok Token) String() string {
	return fmt.Sprintf("%v:%q", tok.Type, tok.Literal)
}

func (tokType TokenType) String() string {
	s := ""
	if 0 <= tokType && tokType < TokenType(len(tokens)) {
		s = tokens[tokType]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tokType)) + ")"
	}
	return s
}

var keywords = map[string]TokenType{
	"package":  PACKAGE,
	"import":   IMPORT,
	"var":      VAR,
	"func":     FUNC,
	"return":   RETURN,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"defer":    DEFER,
	"goto":     GOTO,
}

func Lookup(ident string) TokenType {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}

func (op TokenType) Precedence() int {
	switch op {
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 1
	case ADD, SUB:
		return 2
	case MUL, DIV, MOD:
		return 3
	}
	return 0
}

func (i Token) IntValue() int {
	x, err := strconv.ParseInt(i.Literal, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(x)
}
