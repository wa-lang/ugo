package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/lexer"
	"github.com/chai2010/ugo/token"
)

type Option struct {
	Debug bool
}

func ParseFile(filename string, src []byte, opt Option) (*ast.File, error) {
	p := &parser{
		filename: filename,
		src:      src,
		opt:      opt,
	}
	p.parseFile()
	return p.file, p.err
}

func ParseExpr(src []byte, opt Option) (ast.Expr, error) {
	p := &parser{
		filename: "expr",
		src:      src,
		opt:      opt,
	}
	p.parseExpr()
	return p.expr, p.err
}

type parser struct {
	opt      Option
	filename string
	src      []byte

	input []lexer.Item // the tokens being parsed.
	start int          // start position of this item.
	pos   int          // current position in the input.
	width int          // width of last rune read from input.

	file *ast.File
	node ast.Node
	expr ast.Expr
	err  error
}

func (p *parser) next() lexer.Item {
	if p.pos >= len(p.input) {
		p.width = 0
		return lexer.Item{Token: token.EOF}
	}
	tok := p.input[p.pos]
	p.width = 1
	p.pos += p.width
	return tok
}

func (p *parser) peek() lexer.Item {
	tok := p.next()
	p.backup()
	return tok
}

func (p *parser) backup() {
	p.pos -= p.width
}

func (p *parser) ignore() {
	p.start = p.pos
}

func (p *parser) accept(validTokens []token.Token) bool {
	tok := p.next()
	for _, x := range validTokens {
		if tok.Token == x {
			return true
		}
	}
	p.backup()
	return false
}

func (p *parser) acceptRun(validTokens []token.Token) {
	for p.accept(validTokens) {
	}
	p.backup()
}

func (p *parser) errorf(format string, args ...interface{}) {
	//l.items = append(l.items, Item{
	//	Token:   token.ILLEGAL,
	//	Literal: fmt.Sprintf(format, args...),
	//	Pos:     token.Pos(l.start + 1),
	//})
}
