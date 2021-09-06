package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	add1 []int
	add2 []int

	sum []int
}

func TestMergeTwoLists(t *testing.T) {

	questions := []question{
		{
			add1: []int{1, 2, 3, 4},
			add2: []int{0, 0, 0, 5},
			sum:  []int{0, 0, 0, 1, 2, 3, 4, 5},
		},
		// add unit test cases
		{
			add1: []int{},
			add2: []int{0, 0, 0, 5},
			sum:  []int{0, 0, 0, 5},
		},
		{
			add1: []int{},
			add2: []int{},
			sum:  []int{},
		},
		{
			add1: []int{},
			add2: []int{0},
			sum:  []int{0},
		},
		{
			add1: []int{1, 2, 4},
			add2: []int{1, 3, 4},
			sum:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			add1: []int{-9, 3},
			add2: []int{5, 7},
			sum:  []int{-9, 3, 5, 7},
		},
		{
			add1: []int{-2, -1, 0},
			add2: []int{-3, 0, 1},
			sum:  []int{-3, -2, -1, 0, 0, 1},
		},
	}

	for index, question := range questions {
		fmt.Printf("------ Unit Test Case Index: %d ------\n", index)
		h1 := createList(question.add1)
		h2 := createList(question.add2)

		merge := mergeTwoLists(h1, h2)
		merge.printAll()
	}
}

func createList(nums []int) *ListNode {
	l1 := &ListNode{}
	head1 := l1
	for _, ele := range nums {
		newNode := &ListNode{
			Val:  ele,
			Next: nil,
		}

		newNode.Next = l1.Next
		l1.Next = newNode

		l1 = l1.Next
	}
	head1 = head1.Next
	return head1
}
