package main

/*
LC#1373 二叉树中结点和最大的的二叉查找树
https://leetcode.com/problems/maximum-sum-bst-in-binary-tree/
思路：递归计算左右子树“是否二叉查找树”、“和”，“最小值”、“最大值”，如果算上本身也是二叉查找树且和大于已记录的最大和，则更新
*/
func maxSumBST(root *TreeNode) int {
	result := 0

	var _maxSumBST func(node *TreeNode) (bool, int, int, int)
	_maxSumBST = func(node *TreeNode) (bool, int, int, int) {
		if node == nil {
			return true, 0, 0, 0
		}

		leftBst, leftSum, leftMin, leftMax := _maxSumBST(node.Left)
		rightBst, rightSum, rightMin, rightMax := _maxSumBST(node.Right)

		if node.Left == nil {
			leftMin, leftMax = node.Val, node.Val
		}

		if node.Right == nil {
			rightMin, rightMax = node.Val, node.Val
		}

		if (node.Left == nil || (node.Left != nil && leftMax < node.Val)) &&
			(node.Right == nil || (node.Right != nil && rightMin > node.Val)) &&
			leftBst && rightBst {
			sum := leftSum + rightSum + node.Val
			if sum > result {
				result = sum
			}

			return true, sum, leftMin, rightMax
		}

		return false, 0, 0, 0
	}
	_maxSumBST(root)

	return result
}
