package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var flagDebug = flag.Bool("d", false, "set debug mode")

func main() {
	flag.Parse()
	result := run(exampleExpr)
	fmt.Println(result)
}

func run(node *ExprNode) int {
	if !*flagDebug {
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

//   +
//  / \
// 1   *
//    / \
//   2   +
//      / \
//     3   4
var exampleExpr = &ExprNode{
	Value: "+",
	Left: &ExprNode{
		Value: "1",
	},
	Right: &ExprNode{
		Value: "*",
		Left: &ExprNode{
			Value: "2",
		},
		Right: &ExprNode{
			Value: "+",
			Left: &ExprNode{
				Value: "3",
			},
			Right: &ExprNode{
				Value: "4",
			},
		},
	},
}
