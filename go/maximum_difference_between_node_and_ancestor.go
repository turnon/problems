package main

/**
LC#1026 https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
结点与祖先结点的最大差异
思路：遍历过程中，记录所经历的结点，并用当前结点与它们对比，更新最大差异。论坛解法只传递所经历的最大结点与最小结点，更高效
**/
func maxAncestorDiff(root *TreeNode) int {
	maxDiff := 0

	var recursive func(node *TreeNode, path []int) []int
	recursive = func(node *TreeNode, path []int) []int {
		for _, v := range path {
			diff := node.Val - v
			if diff < 0 {
				diff = -diff
			}
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		path = append(path, node.Val)
		if node.Left != nil {
			oLen := len(path)
			path = recursive(node.Left, path)
			path = path[:oLen]
		}
		if node.Right != nil {
			path = recursive(node.Right, path)
		}
		return path
	}
	recursive(root, []int{})

	return maxDiff
}
