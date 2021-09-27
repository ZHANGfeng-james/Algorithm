package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0002. 两数相加\n"
)

func init() {
	fmt.Println(topic)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// [2,4,3] + [5,6,4]
	// [0] + [0]
	// [0] + [1,2,3]
	// [9,9,9,9,9,9,9] + [9,9,9,9]
	var head *ListNode = &ListNode{}
	sentinel := head

	flag := false
	for l1 != nil || l2 != nil {
		sum, tmp := getSumNode(l1, l2, flag)
		flag = tmp
		newNode := &ListNode{
			Val: sum,
		}
		head.Next = newNode
		head = head.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if flag {
		// add New ListNode
		newNode := &ListNode{
			Val: 1,
		}
		head.Next = newNode
	}

	return sentinel.Next
}

// getSumNode adds two ListNode, and return sum with a flag
func getSumNode(l1 *ListNode, l2 *ListNode, flag bool) (sum int, carry bool) {
	if l1 == nil {
		sum = l2.Val
	}

	if l2 == nil {
		sum = l1.Val
	}

	if l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val
	}

	if flag {
		sum += 1
	}

	if sum > 9 {
		return sum - 10, true
	}
	return sum, false
}
