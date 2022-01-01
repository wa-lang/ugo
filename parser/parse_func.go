package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

func (p *Parser) parseFunc() *ast.FuncDecl {
	// func main()
	tokFunc := p.MustAcceptToken(token.FUNC)
	tokFuncIdent := p.MustAcceptToken(token.IDENT)
	p.MustAcceptToken(token.LPAREN) // (
	p.MustAcceptToken(token.RPAREN) // )

	funcName := &ast.Ident{
		NamePos: tokFuncIdent.Pos,
		Name:    tokFuncIdent.Literal,
	}

	funcType := &ast.FuncType{
		Func: tokFunc.Pos,
	}

	return &ast.FuncDecl{
		Name: funcName,
		Type: funcType,
		Body: p.parseStmt_block(), // {}
	}
}
