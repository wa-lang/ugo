package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

// var x int
// var x int = 2

func (p *parser) parseVar() *ast.VarSpec {
	tokVar := p.r.MustAcceptToken(token.VAR)
	tokIdent := p.r.MustAcceptToken(token.IDENT)

	var varSpec = &ast.VarSpec{
		VarPos: tokVar.Pos,
	}

	varSpec.Name = &ast.Ident{
		NamePos: tokIdent.Pos,
		Name:    tokIdent.IdentName(),
	}

	switch p.r.PeekToken().Type {
	case token.IDENT:
	case token.LBRACK: // []T
	case token.STRUCT:
	case token.MAP:
	case token.INTERFACE:
	default:
	}

	if typ, ok := p.r.AcceptToken(token.IDENT); ok {
		varSpec.Type = &ast.Ident{
			NamePos: typ.Pos,
			Name:    typ.IdentName(),
		}
	}

	if _, ok := p.r.AcceptToken(token.ASSIGN); ok {
		varSpec.Value = p.parseExpr()
	}

	p.r.AcceptTokenList(token.SEMICOLON)
	return varSpec
}
