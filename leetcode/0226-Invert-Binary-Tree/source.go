package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0226. 翻转二叉树\n"
)

func init() {
	fmt.Println(topic)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	// 归并和快排是两种不同的算法思路

	// 归并：自底向上，最底部的数据排序好后，依次向上 merge
	// 快排：自顶向下，由 pivot 将数组拆分成 2 个部分，然后分别在 2 个部分中递归上述步骤

	// 翻转二叉树，应该自底向上执行

	// 递：想要交换 2 和 7 子树，发现 2 和 7 节点各自可看成一棵树，将注意力放在节点 2 和 7 的子树上
	// 归：2/7 均的子树完成了 left 和 right 交换后，完成 4 这颗树的交换

	// head := root
	head := exchangeTreeNode(root)

	// invertTree(root.Left)
	// invertTree(root.Right)

	// // 赋值表达式的右侧是递，左侧是归
	// left := exchangeTreeNode(root.Left)
	// right := exchangeTreeNode(root.Right)

	return head
}

func exchangeTreeNode(node *TreeNode) *TreeNode {
	if node == nil { // 递归终止条件
		return nil
	}
	left := exchangeTreeNode(node.Left)
	right := exchangeTreeNode(node.Right)

	node.Left = right
	node.Right = left

	return node
}
