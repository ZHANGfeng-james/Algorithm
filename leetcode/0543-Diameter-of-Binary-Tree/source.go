package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0543. 二叉树的直径\n"
)

func init() {
	fmt.Println(topic)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Max = 0

func diameterOfBinaryTree(root *TreeNode) int {
	Max = 0
	depth(root)
	return Max
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := depth(root.Left)
	rDepth := depth(root.Right)

	if lDepth+rDepth > Max { // Max 就是路径最大值
		Max = lDepth + rDepth
	}

	return max(lDepth, rDepth) + 1 // 返回的是 root 节点的深度，比如：叶子节点深度为 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
