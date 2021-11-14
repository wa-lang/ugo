package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/logger"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseFile() {
	logger.Debugln("parseFile: peek =", p.peekToken())

	p.file = &ast.File{
		Name: p.filename,
		Data: []byte(p.src),
	}

	// package xxx
	p.parsePackage()

	logger.Debugln("parseFile", 22)

LoopImport:
	for {
		switch tok := p.peekToken(); tok.Type {
		case token.EOF:
			return
		case token.SEMICOLON:
			continue
		case token.IMPORT:
			// import "abc"
			// import xx "abc"
			p.parseImport()
		default:
			logger.Debugln("parseFile", 22, 344)
			break LoopImport
		}
	}

	logger.Debugln("parseFile", 33)

	for {
		switch tok := p.peekToken(); tok.Type {
		case token.EOF:
			return
		case token.SEMICOLON:
			continue
		case token.CONST:

			logger.Debugln("parseFile", 33, 11)
			p.parseConst()
		case token.TYPE:
			logger.Debugln("parseFile", 33, 22)
			p.parseType()
		case token.VAR:
			logger.Debugln("parseFile", 33, 33)
			p.parseVar()
		case token.FUNC:
			logger.Debugln("parseFile", 33, 44)
			p.parseFunc()
		default:
			logger.Debugln("parseFile", 33, 555, tok)
			p.errorf("unknown token: %v", tok)
		}
	}
}

func (p *parser) parsePackage() {
	logger.Debugln("parsePackage: peek =", p.peekToken())

	logger.Debugln(11)
	pkg, ok := p.acceptToken(token.PACKAGE)
	if !ok {
		p.errorf("missing package")
	}

	logger.Debugln(22)
	ident, ok := p.acceptToken(token.IDENT)
	if !ok {
		p.errorf("missing package name")
	}

	logger.Debugln(33)

	p.file.Pkg = &ast.PackageSpec{}

	p.file.Pkg.PkgPos = pkg.Pos
	p.file.Pkg.PkgName = &ast.Ident{
		Name:    ident.IdentName(),
		NamePos: ident.Pos,
	}

	logger.Debugln(44)
}
