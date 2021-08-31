package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	nums   []int
	target int
}

type ans struct {
	one []int
}

func TestProblem(t *testing.T) {
	qs := []question{
		{
			para{[]int{3, 2, 4}, 6},
			ans{[]int{2, 3}},
		},
		{
			para{[]int{3, 2, 4}, 5},
			ans{[]int{1, 2}},
		},
		{
			para{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans{[]int{2, 4}},
		},
		{
			para{[]int{0, 1}, 1},
			ans{[]int{1, 2}},
		},
		{
			para{[]int{0, 3}, 5},
			ans{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("--------Leetcode Problem 0167-------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("[input]:%v;[got]:%v;[want]:%v\n\n", p, twoSum(p.nums, p.target), q.ans)
	}
	fmt.Printf("\n")
}
