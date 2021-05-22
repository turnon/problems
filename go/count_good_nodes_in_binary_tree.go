package main

/**
LC#1448 https://leetcode.com/problems/count-good-nodes-in-binary-tree/
二叉树中的“大于等于所有祖先的”结点的数量
思路：递归所有结点，如果当前结点大于调用栈中最大值，则计数加一，并且更新调用栈中最大值
**/
func goodNodes(root *TreeNode) int {
	count := 0

	var recursive func(node *TreeNode, max int)
	recursive = func(node *TreeNode, max int) {
		if node.Val >= max {
			count += 1
			max = node.Val
		}
		if node.Left != nil {
			recursive(node.Left, max)
		}
		if node.Right != nil {
			recursive(node.Right, max)
		}
	}
	recursive(root, root.Val)

	return count
}
