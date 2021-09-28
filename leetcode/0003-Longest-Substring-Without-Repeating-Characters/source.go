package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0003. 无重复字符的最长子串\n"
)

func init() {
	fmt.Println(topic)
}

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	len := len(s)

	hashmap := make(map[byte]int)

	result := 0
	for end != len {
		ch := s[end]
		if index, ok := hashmap[ch]; !ok {
			// hashmap: ch <--> index
			if result < (end - start + 1) {
				result = end - start + 1
			}
		} else {
			// start -- index 之间的内容删除
			for i := start; i <= index; i++ {
				delete(hashmap, s[i])
			}
			start = index + 1
		}
		hashmap[ch] = end
		end++
	}

	return result
}
