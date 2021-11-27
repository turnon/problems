package main

/*
LC#687 二叉树中最长的同值路径
https://leetcode.com/problems/longest-univalue-path
思路：分别计算左结点和右结点的最长同值路径长度（不含左加右长度），
并在过程中更新树的最长同值路径长度为左长度、右长度、左加右长度中最长的那个
*/
func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	if root == nil {
		return maxLen
	}

	var _longestUnivaluePath func(node *TreeNode) (maxInLeftAndRight int)
	_longestUnivaluePath = func(node *TreeNode) (maxInLeftAndRight int) {
		left, right, milr := 0, 0, 0
		if node.Left != nil {
			milr := _longestUnivaluePath(node.Left)
			if node.Left.Val == node.Val {
				left = milr + 1
			}
		}
		if node.Right != nil {
			milr := _longestUnivaluePath(node.Right)
			if node.Right.Val == node.Val {
				right += milr + 1
			}
		}
		if left > right {
			milr = left
		} else {
			milr = right
		}
		if milr > maxLen {
			maxLen = milr
		}
		if leftandRight := left + right; leftandRight > maxLen {
			maxLen = leftandRight
		}
		return milr
	}

	_longestUnivaluePath(root)
	return maxLen
}
