package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0070. 爬楼梯\n"
)

func init() {
	fmt.Println(topic)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createBinaryTree 依据给定的 nums 数组，构造一个二叉树
func createBinaryTree(nums []int) {

}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// [1,null,2]
	result := make([]int, 0)

	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}
