package leetcode

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for index, ele := range nums {
		sub_value := target - ele
		if i, ok := hashmap[sub_value]; ok {
			return []int{i, index}
		}
		hashmap[ele] = index
	}
	return nil
}
