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
	slow := head
	fast := head

	// [] -1
	// [1] -1
	// [1 2] 0
	// [3 2 0 -4] 1
	// [-21,10,17,8,4,26,5,35,33,-7,-16,27,-12,6,29,-12,5,9,20,14,14,2,13,-24,21,23,-21,5] -1
	for slow != nil && fast != nil {
		slow = slow.Next
		if slow == nil {
			// slow == nil 则 break，不会和 slow == fast 冲突
			break
		}

		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func hasCycleHashmap(head *ListNode) bool {
	hashmap := make(map[*ListNode]struct{})
	pointer := head
	for pointer != nil {
		if _, ok := hashmap[pointer]; ok {
			return true
		}
		hashmap[pointer] = struct{}{}
		pointer = pointer.Next
	}
	return false
}
