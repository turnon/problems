package main

import (
	"strconv"
)

/*
LC#655 把二叉树打印成矩阵
https://leetcode.com/problems/print-binary-tree
思路：本打算用递归，但考虑到可能形成大量数组碎片，改用层序遍历，记录每层结点（包括空结点），然后填充每层的间隔。
可发现每个结点都位于以其为根结点的矩阵的首层的中点，因此可先构造出一层的每个结点的矩阵首层，再将这些首层合并，即构成填充好的一层
*/
func printTree(root *TreeNode) [][]string {
	if root == nil {
		return [][]string{}
	}

	rawLevels := [][]*TreeNode{{root}}
	lastLevel := rawLevels[0]
	for {
		nextLevel := make([]*TreeNode, 0, len(lastLevel)*2)
		nextLevelNodeCount := 0
		for _, n := range lastLevel {
			if n == nil {
				nextLevel = append(nextLevel, nil, nil)
			} else {
				if n.Left != nil {
					nextLevelNodeCount += 1
				}
				if n.Right != nil {
					nextLevelNodeCount += 1
				}
				nextLevel = append(nextLevel, n.Left, n.Right)
			}
		}
		if nextLevelNodeCount == 0 {
			break
		}
		rawLevels = append(rawLevels, nextLevel)
		lastLevel = nextLevel
	}

	bottomWide := (len(rawLevels[len(rawLevels)-1]))*2 - 1

	filledLevels := make([][]string, 0, len(rawLevels))
	for _, rawLevel := range rawLevels {
		filledLevel := make([]string, 0, bottomWide+1)
		fragmentWide := bottomWide / len(rawLevel)
		for nodeIdx := range rawLevel {
			fragment := make([]string, fragmentWide)
			mid := (fragmentWide / 2)
			node := rawLevel[nodeIdx]
			if node != nil {
				fragment[mid] = strconv.Itoa(node.Val)
			}
			filledLevel = append(filledLevel, fragment...)
			filledLevel = append(filledLevel, "")
		}
		filledLevel = filledLevel[0:bottomWide]
		filledLevels = append(filledLevels, filledLevel)
	}

	return filledLevels
}
