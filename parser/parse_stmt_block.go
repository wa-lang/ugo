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

func (p *parser) parseStmt_block() *ast.BlockStmt {
	logger.Debugln("peek =", p.peekToken())

	p.acceptToken(token.LBRACE)

Loop:
	for {
		switch p.peekTokenType() {
		case token.EOF:
			break Loop
		case token.RBRACE:
			break Loop
		default:
		}
	}

	// parse stmt list

	p.acceptToken(token.RBRACE)
	return &ast.BlockStmt{}
}
