# 1 栈

## 1.1 数组实现的栈

一个基础栈数据结构的实现：

~~~go
type stack struct {
	values []rune // Stack的数据容器
	index  int    // 标记栈顶在数据容器的数据索引，是最重要的标识
}

func (stack *stack) push(ele rune) {
	// stack数据结构中去除 capacity 字段
	if stack.index == cap(stack.values)-1 {
		return
	}
	stack.index++
	stack.values[stack.index] = ele
}

func (stack *stack) pop() (rune, bool) {
	if stack.index == -1 {
		return ' ', false
	}

	result := stack.values[stack.index]
	stack.index--
	return result, true
}

func (stack *stack) printAll() {
	// FILO
	for {
		ele, ok := stack.pop()
		if !ok {
			break
		}
		fmt.Printf("ele:%c\n", ele)
	}
}
~~~

测试程序：

~~~go
func testStack() {
	content := "abc"
	len := len(content)
	stack := &stack{
		values:   make([]rune, len),
		index:    -1,
	}

	for _, ele := range content {
		stack.push(ele)
	}

	stack.printAll()
}
~~~

栈——这种**受限的数据结构**——中，**“受限”**主要表现在：push 和 pop 函数的实现上，操作**只能在 index 表示的位置索引上执行**。

其中 index 这个表示是栈顶标识，也就是说：

1. 如果想要 Stack 出栈，直接获取 index 位置的数组元素即可；
2. 如果想要 Stack 进栈，需要在 index + 1 的位置增加入栈的元素。

一个更加标准的写法是这样的：

~~~go
package leetcode

import (
	"fmt"
)

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
		stack.data = append(stack.data, val) // 如果 capacity 足够，则会 reslice，改变 len；反之只会改变 len
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
~~~

## 1.2 链表实现的栈
