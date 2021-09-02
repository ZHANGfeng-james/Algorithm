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
			para{[]int{2, 3, 4}, 6},
			ans{[]int{1, 3}},
		},
		{
			para{[]int{2, 3, 4}, 5},
			ans{[]int{1, 2}},
		},
		{
			para{[]int{0, 2, 3, 5, 7, 8}, 3},
			ans{[]int{1, 3}},
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

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("[input]:%v;[got]:%v;[want]:%v\n\n", p, otherSolution(p.nums, p.target), q.ans)
	}
	fmt.Printf("\n")
}
