package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/logger"
	"github.com/chai2010/ugo/token"
)

// func main() {}
// func name(a int, b int) int
// func (p *Type) method() int
// func() (int, int) {}

func (p *parser) parseFunc() {
	logger.Debugln("peek =", p.peekToken())

	tokFunc, ok := p.acceptToken(token.FUNC)
	if !ok {
		return
	}

	var funcSpec = ast.Func{
		FuncPos: tokFunc.Pos,
	}

	switch p.peekTokenType() {
	case token.IDENT:
		p.parseFunc_func(&funcSpec)
	case token.LPAREN:
		p.parseFunc_method(&funcSpec)
	default:
		p.errorf("invalid token = %v", token.IDENT, p.peekToken())
	}
}

func (p *parser) parseFunc_func(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

	p.parseFunc_name(fn)
	p.parseFunc_args(fn)
	p.parseFunc_returns(fn)
	p.parseFunc_body(fn)

	p.acceptToken(token.SEMICOLON)
}

func (p *parser) parseFunc_method(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}
func (p *parser) parseFunc_closure(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}

func (p *parser) parseFunc_self(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}

func (p *parser) parseFunc_name(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}

func (p *parser) parseFunc_args(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}

func (p *parser) parseFunc_returns(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}

func (p *parser) parseFunc_body(fn *ast.Func) {
	logger.Debugln("peek =", p.peekToken())

}

/*

type Func struct {
	FuncPos token.Pos  // var 关键字位置
	Name    *Ident     // 变量名字
	Args    []*VarSpec // 函数参数
	Returns []*VarSpec // 返回值列表
	Body    *BlockStmt // 函数体
}
*/
