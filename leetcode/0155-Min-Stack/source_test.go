package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println("min:", stack.GetMin())

	stack.Pop()
	fmt.Println("min:", stack.GetMin())

	stack.Pop()
	fmt.Println("min:", stack.GetMin())
}
