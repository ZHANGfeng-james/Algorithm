package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) printAll() {
	for list != nil {
		fmt.Printf("node:%d\n", list.Val)
		list = list.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		// “递”，本次取得 l1 作为当前处理的结果（“归”），取 l1
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		// l1.Val >= l2.Val，取 l2
		// “递”，本次取得 l2 作为当前处理的结果（“归”），取 l2
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
