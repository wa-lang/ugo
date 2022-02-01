package main

import "testing"

func TestRun(t *testing.T) {
	for i, tt := range tests {
		if got := run(tt.code); got != tt.value {
			t.Fatalf("%d: expect = %v, got = %v", i, tt.value, got)
		}
	}
}

var tests = []struct {
	code  *ExprNode
	value int
}{
	{
		code:  exampleExpr,
		value: 15,
	},
}
