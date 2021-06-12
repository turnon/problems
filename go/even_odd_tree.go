package main

/**
LC#1609 https://leetcode.com/problems/even-odd-tree/
奇偶树
思路：层序遍历。需关注当前层的起始和结束，以及奇数层还是偶数层
**/
func isEvenOddTree(root *TreeNode) bool {
	levels := []*TreeNode{root}
	levelEnd := -1
	currentIndex := -1
	oddLevel := true

	addChildren := func(node *TreeNode) {
		if node.Left != nil {
			levels = append(levels, node.Left)
		}
		if node.Right != nil {
			levels = append(levels, node.Right)
		}
	}

	evenOddNotMatch := func(node *TreeNode) bool {
		oddValue := (node.Val%2 == 1)
		if (oddLevel && oddValue) || (!oddLevel && !oddValue) {
			return true
		}
		return false
	}

	for {
		nextEnd := len(levels) - 1
		if levelEnd >= nextEnd {
			break
		}
		levelEnd = nextEnd
		currentIndex++
		oddLevel = !oddLevel

		for ; currentIndex < levelEnd; currentIndex++ {
			currentNode := levels[currentIndex]
			if evenOddNotMatch(currentNode) {
				return false
			}

			if (oddLevel && (currentNode.Val <= levels[currentIndex+1].Val)) ||
				(!oddLevel && (currentNode.Val >= levels[currentIndex+1].Val)) {
				return false
			}

			addChildren(currentNode)
		}

		currentNode := levels[currentIndex]
		if evenOddNotMatch(currentNode) {
			return false
		}
		addChildren(currentNode)
	}

	return true
}
