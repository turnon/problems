package main

/*
LC#669 修剪二叉查找树
https://leetcode.com/problems/trim-a-binary-search-tree
思路：如果当前结点的值小于给定范围，则只修剪右子树并返回右子树；
如果当前结点的值大于给定范围，则只修剪左子树并返回左子树；
如果当前结点的值符合范围，则修剪左右子树并返回当前结点
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if high < root.Val {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
