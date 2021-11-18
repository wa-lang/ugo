package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseStmt_defer(block *ast.BlockStmt) {
	tokDefer := p.mustAcceptToken(token.DEFER)
	callExpr := p.parseExpr_call()

	block.List = append(block.List, &ast.DeferStmt{
		Defer: tokDefer,
		Call:  callExpr,
	})
}
