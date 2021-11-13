package runner

import (
	"os"
	"os/exec"

	"github.com/chai2010/ugo/compiler"
	"github.com/chai2010/ugo/parser"
)

func RunExpr(code []byte) int {
	expr, _ := parser.ParseExpr(code, parser.Option{})
	output := new(compiler.Compiler).CompileExpr(expr)

	os.WriteFile("a.out.ll", []byte(output), 0666)
	err := exec.Command("clang", "-Wno-override-module", "-o", "a.out", "a.out.ll").Run()
	if err != nil {
		return err.(*exec.ExitError).ExitCode()
	}
	return 0
}
