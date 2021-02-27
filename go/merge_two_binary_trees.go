package main

/**
 * LC#617 合并二叉树中相同位置的值
 * 思路：递归，当两结点参数都存在，才合并；否则直接返回另一结点
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		root1.Val + root2.Val,
		mergeTrees(root1.Left, root2.Left),
		mergeTrees(root1.Right, root2.Right),
	}
}
