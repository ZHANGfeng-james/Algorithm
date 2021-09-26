package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0448. 找到所有数组中消失的数字\n"
)

func init() {
	fmt.Println(topic)
}

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	if nums == nil {
		return result
	}

	// range: 1 ~ len 找出在 range 中没有出现在 nums 中的元素
	/*
		输入：nums = [4,3,2,7,8,2,3,1]
		输出：[5,6]
	*/

	eles := make([]bool, len(nums))
	for _, value := range nums {
		// 只有遍历结束后，才能得到最终结果
		eles[value-1] = true
	}

	for index, ele := range eles {
		if !ele {
			result = append(result, index+1)
		}
	}

	return result
}
