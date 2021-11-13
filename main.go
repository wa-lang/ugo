package main

import (
	"io"
	"os"

	"github.com/chai2010/ugo/runner"
)

func main() {
	code, _ := io.ReadAll(os.Stdin)
	runner.RunExpr(code)
}
