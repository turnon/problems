package main

/**
LC#993 https://leetcode.com/problems/cousins-in-binary-tree
表兄弟结点
思路：遍历时记录匹配结点的父结点以及深度，最后如果深度相同而父结点不同，则是表兄弟
**/
func isCousins(root *TreeNode, x int, y int) bool {
	xParent, yParent, xDepth, yDepth := -1, -2, -3, -4

	var recursive func(node *TreeNode, parent int, depth int)
	recursive = func(node *TreeNode, parent int, depth int) {
		if xParent != -1 && yParent != -2 {
			return
		}
		if node.Val == x {
			xParent = parent
			xDepth = depth
		}
		if node.Val == y {
			yParent = parent
			yDepth = depth
		}
		if node.Left != nil {
			recursive(node.Left, node.Val, depth+1)
		}
		if node.Right != nil {
			recursive(node.Right, node.Val, depth+1)
		}
	}
	recursive(root, -1, 0)

	if xParent != yParent && xDepth == yDepth {
		return true
	}
	return false
}
