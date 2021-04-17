package main

/**
LC#1302 https://leetcode.com/problems/deepest-leaves-sum/
最深叶子结点之和
思路：递归找出叶子结点，如果深度不小于纪录，则加入最终结果。如果深度超过纪录，则先更新深度，并把和清零。
**/
func deepestLeavesSum(root *TreeNode) int {
	sum := 0
	maxDepth := 0

	var recursive func(*TreeNode, int)
	recursive = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			if depth >= maxDepth {
				if depth > maxDepth {
					maxDepth = depth
					sum = 0
				}
				sum += node.Val
			}
			return
		}

		depth += 1
		if node.Left != nil {
			recursive(node.Left, depth)
		}
		if node.Right != nil {
			recursive(node.Right, depth)
		}
	}
	recursive(root, 1)

	return sum
}
