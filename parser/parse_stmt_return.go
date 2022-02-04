package parser

import (
	"github.com/wa-lang/ugo/ast"
	"github.com/wa-lang/ugo/token"
)

func (p *Parser) parseStmt_return() *ast.ReturnStmt {
	tokReturn := p.MustAcceptToken(token.RETURN)

	retStmt := &ast.ReturnStmt{
		Return: tokReturn.Pos,
	}
	if _, ok := p.AcceptToken(
		token.SEMICOLON, // ;
		token.LBRACE,    // {
		token.RBRACE,    // }
	); !ok {
		retStmt.Result = p.parseExpr()
	} else {
		p.UnreadToken()
	}

	return retStmt
}
