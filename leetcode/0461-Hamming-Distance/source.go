package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0461. 汉明距离\n"
)

func init() {
	fmt.Println(topic)
}

func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 { // 包含了 3 种情况：x 和 y 都不为 0；x 和 y 有一个不为零（2种情况）
		if x&0x01 != y&0x01 {
			count++
		}
		x = x >> 1
		y = y >> 1
	}
	return count
}
