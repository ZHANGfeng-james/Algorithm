package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para []int
	ans  bool
}

func TestProblem(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[:1])

	slice = slice[:2]
	fmt.Println(slice)

	fmt.Printf("\n")
}
