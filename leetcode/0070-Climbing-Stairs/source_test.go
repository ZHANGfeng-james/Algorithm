package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	step   int
	result int
}

func TestProblem(t *testing.T) {
	qs := []question{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
		{
			5,
			8,
		},
		// 如需多个测试，可以复制上方元素。
		{
			44,
			1134903170,
		},
	}

	for _, q := range qs {
		fmt.Printf("[input]:%v;[got]:%v;[want]:%v\n\n", q.step, climbStairs(q.step), q.result)
	}
	fmt.Printf("\n")
}
