package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chai2010/ugo/lexer"
)

func main() {
	code, _ := io.ReadAll(os.Stdin)

	tokens := lexer.Lex("a.go", string(code), false)
	for _, t := range tokens {
		fmt.Printf("%+v\n", t)
	}
}
