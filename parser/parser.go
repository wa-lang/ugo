package parser

import (
	"github.com/chai2010/ugo/ast"
)

type Option struct {
	Debug bool
}

func ParseFile(name string, src interface{}, opt *Option) (*ast.File, error) {
	panic("TODO")
}

type parser struct{}

func (p *parser) parsePackage() {}
func (p *parser) parseImport()  {}
func (p *parser) parseType()    {}
func (p *parser) parseConst()   {}
func (p *parser) parseVar()     {}
func (p *parser) parseFunc()    {}
