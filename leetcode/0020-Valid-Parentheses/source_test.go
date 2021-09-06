package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

type question struct {
	value  string
	result bool
}

func TestValidParentheses(t *testing.T) {

	questions := []question{
		{
			value:  "{}()",
			result: true,
		},
		{
			value:  "()",
			result: true,
		},
		{
			value:  "()[](){}",
			result: true,
		},
		{
			value:  "(]",
			result: false,
		},
		{
			value:  "({})",
			result: true,
		},
		{
			value:  "",
			result: true,
		},
		{
			value:  "((",
			result: false,
		},
	}

	for _, q := range questions {
		fmt.Printf("question: %v, want: %v, got: %v\n", q.value, q.result, isValid(q.value))
	}

	randomTestCase()
}

func randomTestCase() {
	allChars := "([{)]}"
	length := len(allChars)

	N := 30
	LENGTH := 10
	for index := 0; index < N; index++ {
		randomLen := rand.Intn(LENGTH)
		str := make([]byte, 0)
		for i := 0; i < randomLen; i++ {
			randomIndex := rand.Intn(length)
			str = append(str, allChars[randomIndex])
		}

		fmt.Printf("length:%d, str:%s\n", randomLen, string(str)) // [} ) ( ) [ { (]
	}
}
