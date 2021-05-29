package main

/**
LC#1457 https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
有多少条树根到叶子的路径可转成回文
思路，遍历期间统计路径上的结点数和每个值的重复数。到达叶子时，如果结点数为偶数且没有重复数为奇数，或者结点数为奇数且只有一个重复数为奇数，则路径数加一。
（判断奇偶可用位操作。统计重复数时如果基数已知，则可用数组而不用哈希）
**/
func pseudoPalindromicPaths(root *TreeNode) int {
	countVal := make([]int, 9)
	countNode := 0
	countPath := 0

	canBePalindromicPath := func() bool {
		// 如果结点数为偶数
		if countNode>>1<<1 == countNode {
			for _, vCount := range countVal {
				// 且有任何一个值被收集了奇数次
				if vCount>>1<<1 != vCount {
					return false
				}
			}
			return true
		}

		// 如果结点数为奇数
		oddCount := 0
		for _, vCount := range countVal {
			// 且有多于一个值被收集了奇数次
			if vCount>>1<<1 != vCount {
				oddCount += 1
				if oddCount > 1 {
					return false
				}
			}
		}
		return true
	}

	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		val := node.Val - 1
		countVal[val] = countVal[val] + 1
		countNode += 1

		if node.Left == nil && node.Right == nil {
			if canBePalindromicPath() {
				countPath += 1
			}
		} else {
			if node.Left != nil {
				recursive(node.Left)
			}

			if node.Right != nil {
				recursive(node.Right)
			}
		}

		countVal[val] = countVal[val] - 1
		countNode -= 1
	}
	recursive(root)

	return countPath
}
