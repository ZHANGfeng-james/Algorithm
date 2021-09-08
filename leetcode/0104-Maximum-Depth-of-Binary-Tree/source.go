package leetcode

import (
	"fmt"
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
	return loopTest(root.Left, root.Right)
}

// 封装一个递归判断的函数，必须是同一height
func loopTest(node1, node2 *TreeNode) bool {
	// 第一层，判断 node1 和 node2 的 Val
	if (node1 != nil && node2 == nil) ||
		(node1 == nil && node2 != nil) {
		return false
	}

	if node1 == nil && node2 == nil {
		return true
	}

	if node1.Val != node2.Val {
		return false
	}

	// 第二层判断 left and right tree
	// root 的 left tree's left node 和 right tree's right node
	// root 的 left tree's right node 和 right tree's left node
	return loopTest(node1.Left, node2.Right) && loopTest(node1.Right, node2.Left)
}
