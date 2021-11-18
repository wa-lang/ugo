package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseStmt_if(blk *ast.BlockStmt) {
	tokIf := p.r.MustAcceptToken(token.IF)
	_ = tokIf
}
