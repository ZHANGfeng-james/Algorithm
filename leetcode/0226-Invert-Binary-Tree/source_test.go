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
	nums []int
}

type ans struct {
	one int
}

func TestProblem(t *testing.T) {
	questions := []question{
		{
			para{[]int{1, 2, 2, 2, 3}},
			ans{2},
		},
		{
			para{[]int{3, 2, 3}},
			ans{3},
		},
		{
			para{[]int{2, 2, 1, 1, 1, 2, 2}},
			ans{2},
		},
		{
			para{[]int{1, 1}},
			ans{1},
		},
		{
			para{[]int{1}},
			ans{1},
		},
	}

	for _, q := range questions {
		fmt.Printf("question:%v, got:%d, want:%d\n", q.para, majorityElement(q.para.nums), q.ans)
	}

	fmt.Printf("\n")
}
