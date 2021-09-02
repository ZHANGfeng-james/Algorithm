package leetcode

import "fmt"

const (
	topic = "Leetcode Problem 0167 两数之和 II - 输入有序数组\n"
)

func init() {
	fmt.Println(topic)
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if idx, ok := m[another]; ok {
			return []int{idx + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return nil
}

func otherSolution(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		sub := target - numbers[start]
		if sub == numbers[end] {
			return []int{start + 1, end + 1}
		}

		// [1 2 3 4] 6
		if sub > numbers[end] {
			start++
		} else {
			end--
		}
	}
	return nil
}
