package parser

import (
	"fmt"

	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseExpr() ast.Expr {
	expr := p.parseExpr_mul()
	for {
		switch p.peekToken() {
		case token.ADD, token.SUB:
			tok := p.next()
			expr = &ast.BinaryExpr{
				X:  expr,
				Op: tok.Token,
				Y:  p.parseExpr_mul(),
			}
		default:
			return expr
		}
	}
}

func (p *parser) parseExpr_mul() ast.Expr {
	expr := p.parseExpr_primary()
	for {
		switch p.peekToken() {
		case token.MUL, token.QUO:
			tok := p.next()
			expr = &ast.BinaryExpr{
				X:  expr,
				Op: tok.Token,
				Y:  p.parseExpr_primary(),
			}
		default:
			return expr
		}
	}
}

func (p *parser) parseExpr_primary() ast.Expr {
	switch peek := p.peek(); peek.Token {
	case token.INT:
		p.next()
		switch tok := p.next(); tok.Token {
		case token.INT:
			return &ast.Number{
				Value: int(tok.IntValue()),
			}
		default:
			p.err = fmt.Errorf("todo")
			panic(p.err)
		}
	case token.LPAREN:
		p.next()
		expr := p.parseExpr()
		if !p.accept(token.RPAREN) {
			p.err = fmt.Errorf("todo")
			panic(p.err)
		}
		p.next()
		return expr
	default:
		p.errorf("todo")
		panic(p.err)
	}
}
