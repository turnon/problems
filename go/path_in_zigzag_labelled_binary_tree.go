package main

/**
LC#1104 https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree
在结点顺序为之字形的树中，求某结点的路径
思路：先构造出在非之字形的树中，该结点所在层以及所有上层的每层的起始结点，然后考虑之字形树中，目标结点在最后一层的位置，以及所有父结点在其所属层的位置
**/
func pathInZigZagTree(label int) []int {

	// 非之字形树的最左结点列表
	leftMost := []int{1}
	for i := 1; ; i++ {
		val := 1 << i
		if val > label {
			break
		}
		leftMost = append(leftMost, val)
	}

	// 最后一层中，目标结点的位置
	level := len(leftMost) - 1
	lastLevelIdx := 0
	if level%2 == 0 {
		val := leftMost[level]
		lastLevelIdx = label - val
	} else {
		val := (leftMost[level] << 1) - 1
		lastLevelIdx = val - label
	}
	leftMost[level] = label

	// 计算父结点的位置（当前结点所在层的位置除2）
	for level -= 1; level > 0; level-- {
		lastLevelIdx = lastLevelIdx / 2
		if level%2 == 0 {
			val := leftMost[level]
			leftMost[level] = val + lastLevelIdx
		} else {
			val := (leftMost[level] << 1) - 1
			leftMost[level] = val - lastLevelIdx
		}
	}

	return leftMost
}
