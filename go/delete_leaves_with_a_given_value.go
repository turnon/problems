package main

/**
LC#1325 https://leetcode.com/problems/delete-leaves-with-a-given-value/
使树中没有任何叶子等于指定值
思路：后序遍历，如果左右结点都被删掉，且自身也等于指定值，则自身也要删掉
**/
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}

	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}
