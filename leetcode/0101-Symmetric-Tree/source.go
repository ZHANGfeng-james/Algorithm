package leetcode

import (
	"fmt"
	"math"
)

const (
	topic = "Leetcode Problem 0101. 对称二叉树\n"
)

func init() {
	fmt.Println(topic)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	tmp := inorderTraversal(root)

	length := len(tmp)
	for index, ele := range tmp {
		if ele != tmp[length-index-1] {
			return false
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// [1,2,2,2,null,2]

	result := make([]int, 0)

	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
		return result
	}

	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	} else {
		result = append(result, math.MaxInt32)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	} else {
		result = append(result, math.MaxInt32)
	}
	return result
}
