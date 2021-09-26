package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0338. 比特位计数\n"
)

func init() {
	fmt.Println(topic)
}

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0

		// count Bits
		tmp := i
		for tmp != 0 {
			if (tmp & 0x01) == 1 {
				count++
			}
			tmp = tmp >> 1
		}

		result[i] = count
	}

	return result
}
