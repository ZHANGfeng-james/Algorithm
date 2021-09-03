package leetcode

import (
	"fmt"
)

var rightMap = make(map[rune]bool)
var pairMap = make(map[rune]rune)

func init() {
	rightMap[')'] = true
	rightMap[']'] = true
	rightMap['}'] = true

	pairMap['('] = ')'
	pairMap['['] = ']'
	pairMap['{'] = '}'
}

func isRight(value rune) bool {
	_, ok := rightMap[value]
	return ok
}

func isPair(left, right rune) bool {
	result, ok := pairMap[left]
	if ok {
		if result == right {
			return true
		}
	}
	return false
}

func isValid(content string) bool {
	len := len(content)
	stack := &stack{
		capacity: len,
		values:   make([]rune, len),
		index:    -1,
	}
	return stack.isValidParentheses(content)
}

type stack struct {
	capacity int
	values   []rune
	index    int
}

func (stack *stack) push(ele rune) {
	if stack.index == stack.capacity-1 {
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

func (stack *stack) isValidParentheses(content string) bool {
	if len(content)%2 == 1 {
		return false
	}

	var result = true
	for _, ele := range content {
		if isRight(ele) {
			pop, ok := stack.pop()
			if ok {
				if !isPair(pop, ele) {
					result = false
					break
				} else {
					continue
				}
			} else {
				result = false
				// stack empty
				break
			}
		}
		stack.push(ele)
	}

	if stack.index != -1 {
		result = false
	}

	return result
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

func testStack() {
	content := "abc"
	len := len(content)
	stack := &stack{
		capacity: len,
		values:   make([]rune, len),
		index:    -1,
	}

	for _, ele := range content {
		stack.push(ele)
	}

	stack.printAll()
}
