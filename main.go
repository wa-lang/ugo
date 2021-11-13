package main

import (
	"fmt"
	"io"
	"os"

	_ "github.com/chai2010/ugo/compiler"
	"github.com/chai2010/ugo/lexer"
	_ "github.com/chai2010/ugo/parser"
)

func main() {
	code, _ := io.ReadAll(os.Stdin)

	tokens := lexer.Lex("a.go", string(code), lexer.Option{
		SkipComment:    false,
		DontInsertSemi: true,
	})
	for _, t := range tokens {
		fmt.Printf("%+v\n", t)
	}
}
