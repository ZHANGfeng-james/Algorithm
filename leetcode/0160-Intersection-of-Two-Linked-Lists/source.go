package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0160. 相交链表\n"
)

func init() {
	fmt.Println(topic)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashmap := make(map[*ListNode]bool)
	cursor := headA
	for cursor != nil {
		hashmap[cursor] = true
		cursor = cursor.Next
	}
	cursor = headB
	for cursor != nil {
		if _, ok := hashmap[cursor]; ok {
			break
		}
		cursor = cursor.Next // 如果没有相交的节点，返回值是 nil
	}
	return cursor
}
