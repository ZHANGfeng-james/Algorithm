package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0283. 移动零\n"
)

func init() {
	fmt.Println(topic)
}

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	last := len(nums) - 1
	i, j := last, last
	for i >= 0 {
		if nums[i] == 0 {
			moveEle(nums, i, j)
			j--
		}
		i--
	}
}

func moveEle(nums []int, i, j int) {
	for i != j {
		nums[i] = nums[i+1]
		i++
	}
	nums[j] = 0
}
