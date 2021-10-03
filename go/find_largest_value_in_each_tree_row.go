package main

/*
LC#515 二叉树中每行里最大的数
https://leetcode.com/problems/find-largest-value-in-each-tree-row
思路：层序遍历，过程中记录每层最大的数
*/
func largestValues(root *TreeNode) []int {
	maxValues := []int{}
	if root == nil {
		return maxValues
	}

	level := []*TreeNode{root}
	length := len(level)
	for length > 0 {
		maxValue := level[0].Val
		nextLevel := make([]*TreeNode, 0, 2*length)
		for _, node := range level {
			if node.Val > maxValue {
				maxValue = node.Val
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level = nextLevel
		length = len(level)
		maxValues = append(maxValues, maxValue)
	}

	return maxValues
}
