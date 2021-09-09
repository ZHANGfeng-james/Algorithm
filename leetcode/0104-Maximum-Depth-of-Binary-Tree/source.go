package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0104. 二叉树的最大深度\n"
)

func init() {
	fmt.Println(topic)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// [1,null,2]
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
