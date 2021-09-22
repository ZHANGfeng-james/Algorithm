package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0234. 回文链表\n"
)

func init() {
	fmt.Println(topic)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	topIndex := -1

	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		topIndex++
		cur = cur.Next
	}

	for head != nil {
		// forward and backward
		if head.Val != stack[topIndex] {
			return false
		}

		head = head.Next
		stack = stack[:topIndex]
		topIndex--
	}

	return true
}
