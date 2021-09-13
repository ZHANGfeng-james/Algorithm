package leetcode

import (
	"fmt"
	"math"
)

const (
	topic = "Leetcode Problem 0155. 最小栈\n"
)

func init() {
	fmt.Println(topic)
}

type MinStack struct {
	data []int
	min  []int // 辅助Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		data: []int{},
		min:  []int{math.MaxInt64},
	}
	fmt.Println(len(stack.data), cap(stack.data))
	return stack
}

// Push a element to stack using []int
func (stack *MinStack) Push(val int) {
	stack.data = append(stack.data, val)

	top := stack.min[len(stack.min)-1]
	// int 类型到底占多少字节？默认是 int32 还是 int64？
	stack.min = append(stack.min, min(top, val))
}

func (stack *MinStack) Pop() {
	stack.data = stack.data[:len(stack.data)-1]
	stack.min = stack.min[:len(stack.min)-1]
}

func (stack *MinStack) Top() int {
	return stack.data[len(stack.data)-1]
}

func (stack *MinStack) GetMin() int {
	return stack.min[len(stack.min)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
