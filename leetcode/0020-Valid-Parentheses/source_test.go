package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	value  string
	result bool
}

func TestValidParentheses(t *testing.T) {

	questions := []question{
		{
			value:  "{}()",
			result: true,
		},
		{
			value:  "()",
			result: true,
		},
		{
			value:  "()[](){}",
			result: true,
		},
		{
			value:  "(]",
			result: false,
		},
		{
			value:  "({})",
			result: true,
		},
		{
			value:  "",
			result: true,
		},
		{
			value:  "((",
			result: false,
		},
	}

	for _, q := range questions {
		fmt.Printf("question: %v, want: %v, got: %v\n", q.value, q.result, isValid(q.value))
	}
}
