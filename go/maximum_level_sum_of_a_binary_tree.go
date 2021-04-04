package main

/**
LC#1161 https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
最大的一层之和出现在第几层
思路：层序遍历，构造下一层的同时，对当前层求和，如果和大于以往记录，则更新最大和，以及层数，最后返回层数
**/
func maxLevelSum(root *TreeNode) int {
	max := root.Val
	level := []*TreeNode{root}
	levelN := 0
	levelMax := 1

	for {
		nodeCount := len(level)
		if nodeCount == 0 {
			break
		}
		levelN += 1
		newLevel := make([]*TreeNode, 0, 2*nodeCount)
		sum := 0
		for _, node := range level {
			sum += node.Val
			if node.Left != nil {
				newLevel = append(newLevel, node.Left)
			}
			if node.Right != nil {
				newLevel = append(newLevel, node.Right)
			}
		}
		if max < sum {
			max = sum
			levelMax = levelN
		}
		level = newLevel
	}

	return levelMax
}
