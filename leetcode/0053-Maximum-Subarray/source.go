package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0053. 最大子序和\n"
)

func init() {
	fmt.Println(topic)
}

func maxSubArray(nums []int) int {
	// len(nums) > 0
	result := nums[0]
	for index := 1; index < len(nums); index++ {
		if nums[index]+nums[index-1] > nums[index] {
			nums[index] += nums[index-1]
		}
		if nums[index] > result {
			result = nums[index]
		}
	}
	return result
}

func bruteForce(nums []int) int {
	// []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// Brute Force 方法解决思路：二重循环，遍历所有可能的和值
	// len(nums) > 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
