package main

/**
LC#958 https://leetcode.com/problems/check-completeness-of-a-binary-tree
检查是否完整二叉树
思路：层序遍历，如果遇到空结点后，还有结点，则不是完整二叉树
**/
func isCompleteTree(root *TreeNode) bool {
	level := []*TreeNode{root}

	gotNil := false
	for {
		var nextLevel []*TreeNode
		for _, node := range level {
			if node == nil {
				gotNil = true
				continue
			}

			if gotNil {
				return false
			}

			nextLevel = append(nextLevel, node.Left, node.Right)
		}

		if len(nextLevel) == 0 {
			break
		}
		level = nextLevel
	}

	return true
}
