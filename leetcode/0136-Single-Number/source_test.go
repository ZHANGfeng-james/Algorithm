package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para []int
	ans  int
}

func TestProblem(t *testing.T) {
	qs := []question{
		{
			para: []int{2, 2, 4},
			ans:  4,
		},
		{
			para: []int{1, 3, 1},
			ans:  3,
		},
		{
			para: []int{5, 5, 0},
			ans:  0,
		},
		{
			para: []int{3, 2, 2},
			ans:  3,
		},
		{
			para: []int{0, 0, 1},
			ans:  1,
		},
		// 如需多个测试，可以复制上方元素。
		{
			para: []int{1, 0, 1},
			ans:  0,
		},
	}

	for _, q := range qs {
		fmt.Printf("[input]:%v;[got]:%v;[want]:%v\n\n", q.para, singleNumber(q.para), q.ans)
	}
	fmt.Printf("\n")
}
