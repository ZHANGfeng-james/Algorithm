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

## 1.2 链表实现的栈
