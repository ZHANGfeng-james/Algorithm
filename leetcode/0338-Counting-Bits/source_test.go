package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para int
	ans  []int
}

func TestProblem(t *testing.T) {

	questions := []question{
		{
			0,
			[]int{0},
		},
		// add unit test case
		{
			1,
			[]int{0, 1},
		},
		{
			2,
			[]int{0, 1, 1},
		},
		{
			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, q := range questions {
		fmt.Printf("question:%v, got:%v, want:%v\n", q.para, countBits(q.para), q.ans)
	}
}
