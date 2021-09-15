package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0206. 反转链表\n"
)

func init() {
	fmt.Println(topic)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	cur := prev.Next
	// cur != nil
	next := cur.Next

	prev.Next = nil
	// cur.Next = prev
	// [1,2,3,4,5] --> [5,4,3,2,1]
	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = next
		if cur != nil {
			next = cur.Next
		}
	}

	return prev
}

func reverse(prev *ListNode, cur *ListNode, next *ListNode) {
	// 反转链表，只能从顺次遍历节点时开始
	cur.Next = prev
	// 节点平移
	prev = cur
	cur = next
	next = cur.Next
}
