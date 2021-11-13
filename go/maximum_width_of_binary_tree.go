package main

/*
LC#662 二叉树的宽度
https://leetcode.com/problems/maximum-width-of-binary-tree
思路：遍历过程中带上结点所在层数和层中位置，记录每层的最小位置和最大位置，最后计算每层的宽度
*/
func widthOfBinaryTree(root *TreeNode) int {
	levelMin, levelMax := make(map[int]int), make(map[int]int)

	var walk func(node *TreeNode, level int, idxInLevel int)
	walk = func(node *TreeNode, level int, idxInLevel int) {
		if node.Left != nil {
			walk(node.Left, level+1, idxInLevel<<1)
		}
		if node.Right != nil {
			walk(node.Right, level+1, (idxInLevel<<1)+1)
		}
		if lmin, ok := levelMin[level]; !ok || lmin > idxInLevel {
			levelMin[level] = idxInLevel
		}
		if lmax, ok := levelMax[level]; !ok || lmax < idxInLevel {
			levelMax[level] = idxInLevel
		}

	}
	walk(root, 0, 0)

	maxWid := 0
	for level, min := range levelMin {
		max := levelMax[level]
		width := max - min + 1
		if width > maxWid {
			maxWid = width
		}
	}

	return maxWid
}
