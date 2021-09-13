package leetcode

import (
	"fmt"
)

// https://github.com/wangzheng0822/algo/blob/master/go/08_stack/StackBasedOnArray.go
type ArrayStack struct {
	data []interface{}
	top  int // 栈顶指针，也就是当前slice的一个位置索引值
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 32), // len:0 cap:32，初始容量
		top:  -1,
	}
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.top < 0
}

func (stack *ArrayStack) PopEle() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	v := stack.data[stack.top]
	stack.top--
	return v
}

func (stack *ArrayStack) PushEle(val interface{}) {
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top++
	}

	if stack.top > len(stack.data)-1 {
		stack.data = append(stack.data, val)
	} else {
		stack.data[stack.top] = val
	}
}

func (stack *ArrayStack) PeekEle() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

func (stack *ArrayStack) PrintAll() {
	if stack.IsEmpty() {
		fmt.Println("ArrayStack is Empty!")
	} else {
		for index := stack.top; index >= 0; index-- {
			fmt.Printf("stack.data[%d]= %v\n", index, stack.data[index])
		}
	}
}
