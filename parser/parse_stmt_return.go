package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseStmt_return(block *ast.BlockStmt) {
	tokReturn := p.mustAcceptToken(token.RETURN)
	exprs := p.parseExprList()

	block.List = append(block.List, &ast.ReturnStmt{
		Result:  tokReturn,
		Results: exprs,
	})
}
