package leetcode

import "fmt"

const (
	topic = "\n 【Leetcode Problem 0136 只出现一次的数字】 \n"
)

func init() {
	fmt.Println(topic)
}

type Element struct {
	ele     int
	hasSame bool
}

func bruteForce(nums []int) int {
	// [4,1,2,1,2]
	// [1, 1, 0]
	// [1,0,1]

	// 借助外部容器
	objs := make([]Element, len(nums))
	for i, value := range nums {
		objs[i] = Element{
			ele: value,
		}
	}

	for i, value := range objs {
		if value.hasSame {
			continue
		}

		if i == len(nums)-1 {
			return value.ele
		}

		for j := i + 1; j < len(nums); j++ {
			if objs[j].ele == value.ele {
				objs[j].hasSame = true
				value.hasSame = true
				break
			} else {
				if j == len(nums)-1 {
					return value.ele
				}
			}
		}
	}
	return nums[0]
}

func useHashMap(nums []int) int {
	// key: nums 的元素值; value: false --> true --> false
	hashMap := make(map[int]bool)
	for _, value := range nums {
		if _, ok := hashMap[value]; ok {
			hashMap[value] = !hashMap[value]
		} else {
			hashMap[value] = true
		}
	}

	for key, value := range hashMap {
		if value {
			return key
		}
	}
	return nums[0]
}

func singleNumber(nums []int) int {
	result := 0
	for index := len(nums) - 1; index >= 0; index-- {
		result ^= nums[index]
	}
	return result
}
