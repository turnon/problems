package main

/**
LC#508 https://leetcode.com/problems/most-frequent-subtree-sum
最常见的子树和
思路：后序遍历，得当前树的和，收集到map中，并更新常见度，最后筛选出map中常见度等于最大常见度的和
**/
func findFrequentTreeSum(root *TreeNode) []int {
	freq := map[int]int{}
	maxFreq := 0
	addFreq := func(sum int) {
		count, exists := freq[sum]
		if !exists {
			count = 0
		}
		count += 1
		if count > maxFreq {
			maxFreq = count
		}
		freq[sum] = count
	}

	var _findFrequentTreeSum func(node *TreeNode) int
	_findFrequentTreeSum = func(node *TreeNode) int {
		leftSum, rightSum := 0, 0
		if node.Left != nil {
			leftSum = _findFrequentTreeSum(node.Left)
		}
		if node.Right != nil {
			rightSum = _findFrequentTreeSum(node.Right)
		}
		sum := node.Val + leftSum + rightSum
		addFreq(sum)
		return sum
	}
	_findFrequentTreeSum(root)

	result := make([]int, 0, len(freq))
	for sum, f := range freq {
		if f == maxFreq {
			result = append(result, sum)
		}
	}

	return result
}
