package main

/**
LC#1530 https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/
有多少两叶距离少于指定值
思路：深度优先遍历，左树所有叶子的深度逐一与右树的相加，如果符合，就计数加一。然后左右叶子合并作为本树叶子集合提交上层计算。
**/
func countPairs(root *TreeNode, distance int) int {
	c := 0

	addOneDepth := func(distnces []int) {
		for i, _ := range distnces {
			distnces[i] += 1
		}
	}

	var recursive func(*TreeNode) []int
	recursive = func(node *TreeNode) []int {
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}

		leftLeavesDepths, rightLeavesDepths := []int{}, []int{}
		if node.Left != nil {
			leftLeavesDepths = recursive(node.Left)

		}
		if node.Right != nil {
			rightLeavesDepths = recursive(node.Right)

		}

		addOneDepth(leftLeavesDepths)
		addOneDepth(rightLeavesDepths)

		for _, l := range leftLeavesDepths {
			for _, r := range rightLeavesDepths {
				if l+r <= distance {
					c += 1
				}
			}
		}

		return append(leftLeavesDepths, rightLeavesDepths...)
	}
	recursive(root)

	return c
}
