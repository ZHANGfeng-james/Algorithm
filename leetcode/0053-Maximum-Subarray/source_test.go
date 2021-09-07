package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	nums   []int
	target int
}

func TestProblem(t *testing.T) {
	qs := []question{
		{
			[]int{2, 3, 4},
			9,
		},
		{
			[]int{-2, 3, 4},
			7,
		},
		{
			[]int{2, -3, 4},
			4,
		},
		{
			[]int{4},
			4,
		},
		{
			[]int{2, 3, -4},
			5,
		},
		// 如需多个测试，可以复制上方元素。
		{
			[]int{-1000},
			-1000,
		},
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
		{
			[]int{-1, -2, -3},
			-1,
		},
		{
			[]int{-1, 0, 1},
			1,
		},
		{
			[]int{1, -1, 1},
			1,
		},
	}

	for _, q := range qs {
		fmt.Printf("[input]:%v;[got]:%v;[want]:%v\n\n", q.nums, maxSubArray(q.nums), q.target)
	}
	fmt.Printf("\n")
}
