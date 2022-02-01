package ast

var (
	_ Stmt = (*ExprStmt)(nil)

	_ Expr = (*BinaryExpr)(nil)
	_ Expr = (*CallExpr)(nil)
	_ Expr = (*Number)(nil)
)

func (p *ExprStmt) Pos() int   { return 0 }
func (p *ExprStmt) End() int   { return 0 }
func (p *ExprStmt) stmt_type() {}

func (p *BinaryExpr) Pos() int   { return 0 }
func (p *BinaryExpr) End() int   { return 0 }
func (p *BinaryExpr) expr_type() {}

func (p *CallExpr) Pos() int   { return 0 }
func (p *CallExpr) End() int   { return 0 }
func (p *CallExpr) expr_type() {}

func (p *Number) Pos() int   { return 0 }
func (p *Number) End() int   { return 0 }
func (p *Number) expr_type() {}
