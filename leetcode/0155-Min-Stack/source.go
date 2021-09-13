package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0155. 最小栈\n"
)

func init() {
	fmt.Println(topic)
}

type Element struct {
	data int
	min  int
}

type MinStack struct {
	ele []Element // 每个元素都是一个元组：data, min
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		ele: make([]Element, 0),
	}
	return stack
}

// Push a element to stack using []int
func (stack *MinStack) Push(val int) {
	if len(stack.ele) == 0 {
		stack.ele = append(stack.ele, Element{
			data: val,
			min:  val,
		})
		return
	}

	min := stack.ele[len(stack.ele)-1].min
	stack.ele = append(stack.ele, Element{
		data: val,
		min:  minFunc(val, min),
	})
}

func (stack *MinStack) Pop() {
	stack.ele = stack.ele[:len(stack.ele)-1]
}

func (stack *MinStack) Top() int {
	return stack.ele[len(stack.ele)-1].data
}

func (stack *MinStack) GetMin() int {
	return stack.ele[len(stack.ele)-1].min
}

func minFunc(x, y int) int {
	if x < y {
		return x
	}
	return y
}
