package main

/**
LC#938 https://leetcode.com/problems/range-sum-of-bst/
二叉查找树中符合某区间的结点之和
思路：递归遍历，如果当前结点已在区间以外，则停止下探
**/
func rangeSumBST(root *TreeNode, low int, high int) int {
	curr := 0
	if low <= root.Val && root.Val <= high {
		curr = root.Val
	}

	left := 0
	if root.Left != nil && low <= root.Val {
		left = rangeSumBST(root.Left, low, high)
	}

	right := 0
	if root.Right != nil && root.Val <= high {
		right = rangeSumBST(root.Right, low, high)
	}

	return curr + left + right
}
