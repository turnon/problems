package main

/**
LC#897 https://leetcode.com/problems/increasing-order-search-tree/
二叉查找树转为只有右子树
思路：中序遍历，处理当前结点时生成新结点，挂到上次生成的结点的右结点
**/
func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{0, nil, nil}
	prev := head

	var travelInOrder func(node *TreeNode)

	travelInOrder = func(node *TreeNode) {
		if node.Left != nil {
			travelInOrder(node.Left)
		}
		prev.Right = &TreeNode{node.Val, nil, nil}
		prev = prev.Right
		if node.Right != nil {
			travelInOrder(node.Right)
		}
	}
	travelInOrder(root)
	return head.Right
}
