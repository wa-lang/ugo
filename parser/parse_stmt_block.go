package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/logger"
	"github.com/chai2010/ugo/token"
)

// var x ...
// x := ...
// x, y = ...
// x, y := fn()

// if expr {} else {}
// for {}
// for expr {}
// for stmt; expr; stmt {}
// return expr?

func (p *parser) parseStmt_block() (block *ast.BlockStmt) {
	logger.Debugln("peek =", p.r.PeekToken())

	block = new(ast.BlockStmt)
	p.r.MustAcceptToken(token.LBRACE)

Loop:
	for {
		switch tok := p.r.PeekToken(); tok.Type {
		case token.EOF:
			break Loop
		case token.ILLEGAL:
			panic(tok)
		case token.SEMICOLON:
			p.r.AcceptTokenList(token.SEMICOLON)

		case token.RBRACE:
			break Loop

		case token.CONST:
			_ = p.parseConst()
		case token.TYPE:
			_ = p.parseType()
		case token.VAR:
			_ = p.parseVar()

		case token.DEFER:
			p.parseStmt_defer(block)
		case token.IF:
			p.parseStmt_if(block)
		case token.FOR:
			p.parseStmt_for(block)
		case token.RETURN:
			p.parseStmt_return(block)

		default:
			p.parseStmt_assign(block)
		}
	}

	// parse stmt list

	p.r.AcceptToken(token.RBRACE)
	return
}
