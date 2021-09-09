package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0141 环形链表\n"
)

func init() {
	fmt.Println(topic)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	return false
}
