package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	s   string
	ans int
}

func TestProblem(t *testing.T) {
	questions := []question{
		{"abcabcbb", 3},
		{"aaaaa", 1},
		{"", 0},
		{"pwwkew", 3},
	}

	for _, q := range questions {
		fmt.Printf("question:%s, got:%d, want:%d\n", q.s, lengthOfLongestSubstring(q.s), q.ans)
	}
}
