package leetcode

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {

	stack := NewArrayStack()

	stack.PushEle(1)
	stack.PushEle(2)
	stack.PushEle(3)
	stack.PrintAll()

	fmt.Println("<--->")
	stack.PopEle()
	stack.PrintAll()

	fmt.Println("<--->")
	stack.PushEle(4)
	stack.PrintAll()

	fmt.Println("<--->")
	stack.PopEle()
	stack.PopEle()
	stack.PopEle()
	stack.PrintAll()

}
