package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0070. 爬楼梯\n"
)

func init() {
	fmt.Println(topic)
}

var usage = make(map[int]int)

func init() {
	usage[1] = 1
	usage[2] = 2
	usage[3] = 3
}

func climbStairs(n int) int {
	if result, ok := usage[n]; ok {
		return result
	}

	result := climbStairs(n-1) + climbStairs(n-2)
	usage[n] = result
	return result
}
