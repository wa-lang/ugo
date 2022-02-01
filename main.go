package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var flagDebug = flag.Bool("d", false, "set debug mode")

func main() {
	flag.Parse()

	//   +
	//  / \
	// 1   *
	//    / \
	//   2   +
	//      / \
	//     3   4
	expr_tokens := []string{"1", "+", "2", "*", "(", "3", "+", "4", ")"}

	ast := ParseExpr(expr_tokens)
	fmt.Println(JSONString(ast))

	fmt.Println(run(ast))
}

func run(node *ExprNode) (ret int) {
	if !*flagDebug && ret == 0 {
		defer os.Remove("a.out.ll")
		defer os.Remove("a.out")
	}

	compile(node)
	if err := exec.Command("./a.out").Run(); err != nil {
		return err.(*exec.ExitError).ExitCode()
	}
	return 0
}

func compile(node *ExprNode) {
	output := new(Compiler).GenLLIR(node)

	os.WriteFile("a.out.ll", []byte(output), 0666)
	exec.Command("clang", "-Wno-override-module", "-o", "a.out", "a.out.ll").Run()
}

func JSONString(x interface{}) string {
	d, _ := json.MarshalIndent(x, "", "    ")
	return string(d)
}
