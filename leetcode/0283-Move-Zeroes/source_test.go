package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para []int
	ans  []int
}

func TestProblem(t *testing.T) {

	questions := []question{
		{
			[]int{0, 0, 1, 2, 3},
			[]int{1, 2, 3, 0, 0},
		},
		// add unit test case
		{
			[]int{0, 1, 1, 0, 3},
			[]int{1, 1, 3, 0, 0},
		},
	}

	for _, q := range questions {
		origin := make([]int, len(q.para))
		copy(origin, q.para)

		moveZeroes(q.para)
		fmt.Printf("question:%v, got:%v, want:%v\n", origin, q.para, q.ans)
	}
}
