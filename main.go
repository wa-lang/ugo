package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var flagDebug = flag.Bool("d", false, "set debug mode")

func main() {
	flag.Parse()

	code, _ := io.ReadAll(os.Stdin)
	tokens := Lex(string(code))
	fmt.Println(tokens)

	ast := ParseExpr(tokens)
	fmt.Println(JSONString(ast))

	fmt.Println(run(ast))
}

func run(node *ExprNode) (ret int) {
	if !*flagDebug && ret == 0 {
		defer os.Remove("a.out.ll")
		defer os.Remove("a.out")
	}
	compile(node)
	if data, err := exec.Command("./a.out").CombinedOutput(); err != nil {
		fmt.Print(string(data))
		return err.(*exec.ExitError).ExitCode()
	}
	return 0
}

func compile(node *ExprNode) {
	output := new(Compiler).GenLLIR(node)

	os.WriteFile("a.out.ll", []byte(output), 0666)
	data, err := exec.Command("clang", "-Wno-override-module", "-o", "a.out", "a.out.ll").CombinedOutput()
	if err != nil {
		fmt.Print(string(data))
		os.Exit(1)
	}
}

func JSONString(x interface{}) string {
	d, _ := json.MarshalIndent(x, "", "    ")
	return string(d)
}
