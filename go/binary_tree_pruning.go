package main

/*
LC#814 清除不含1的子树
https://leetcode.com/problems/binary-tree-pruning
思路：后序遍历，告知上层本结点和左右子树是否含1。过程中如果子树不含1，则将其移除，
*/
func pruneTree(root *TreeNode) *TreeNode {
	var hasOne func(node *TreeNode) bool
	hasOne = func(node *TreeNode) bool {
		leftHasOne := false
		if node.Left != nil {
			leftHasOne = hasOne(node.Left)
			if !leftHasOne {
				node.Left = nil
			}
		}

		rightHasOne := false
		if node.Right != nil {
			rightHasOne = hasOne(node.Right)
			if !rightHasOne {
				node.Right = nil
			}
		}

		return node.Val == 1 || leftHasOne || rightHasOne
	}

	if hasOne(root) {
		return root
	}
	return nil
}
