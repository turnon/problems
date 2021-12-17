package main

/*
LC#783 二叉查找树中任意两点的差值的最小值
https://leetcode.com/problems/minimum-distance-between-bst-nodes
思路：中序遍历，记住上一结点。当前结点与上一结点的差值若小于已记录的最小值，则更新。
*/
func minDiffInBST(root *TreeNode) int {
	var min *int

	var prev *TreeNode
	updateMin := func(current *TreeNode) {
		if prev == nil {
			prev = current
			return
		}
		diff := current.Val - prev.Val
		prev = current
		if diff < 0 {
			diff = -diff
		}
		if min == nil || diff < *min {
			min = &diff
		}
	}

	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node.Left != nil {
			walk(node.Left)
		}
		updateMin(node)
		if node.Right != nil {
			walk(node.Right)
		}
	}
	walk(root)

	return *min
}
