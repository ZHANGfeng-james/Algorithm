package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	x   int
	y   int
	ans int
}

func TestProblem(t *testing.T) {

	questions := []question{
		{
			1, 4, 2,
		},
		// add unit test case
		{
			1, 1, 0,
		},
	}

	for _, q := range questions {
		fmt.Printf("question:x=%v, y=%v, got:%v, want:%v\n", q.x, q.y, hammingDistance(q.x, q.y), q.ans)
	}
}
