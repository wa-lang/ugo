package main

import (
	"encoding/json"
	"fmt"

	"github.com/wa-lang/ugo/ast"
	"github.com/wa-lang/ugo/token"
)

func main() {
	fmt.Println(JSONString(ugoProg))
}

func JSONString(x interface{}) string {
	d, _ := json.MarshalIndent(x, "", "    ")
	return string(d)
}

var ugoProg = &ast.File{
	Pkg: &ast.Package{
		Name: "main",
	},
	Funcs: []*ast.Func{
		{
			Name: "main",
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.CallExpr{
							FuncName: "exit",
							Args: []ast.Expr{
								&ast.BinaryExpr{
									Op: token.Token{Type: token.ADD},
									X:  &ast.Number{Value: 40},
									Y:  &ast.Number{Value: 2},
								},
							},
						},
					},
				},
			},
		},
	},
}
