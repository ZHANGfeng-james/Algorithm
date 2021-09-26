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
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		// add unit test case
		{
			[]int{1, 4, 5, 2, 1},
			[]int{3},
		},
	}

	for _, q := range questions {
		fmt.Printf("question:%v, got:%v, want:%v\n", q.para, findDisappearedNumbers(q.para), q.ans)
	}
}
