package main

/**
LC#1372 https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree
二叉树中Z字形路径的最大长度
思路：递归计算每个结点底下的Z字形路径长度（一加左结点的右路径，一加右结点的左路径），并将长度缓存以便上层取到下层任一方向的路径。并在计算时更新最大长度。论坛是返回两个值给上层没有使用缓存。
**/
func longestZigZag(root *TreeNode) int {

	nodeZzCache := make(map[*TreeNode][2]int)
	left := 0
	right := 1
	max := 0

	updateMax := func(length int) {
		if length > max {
			max = length
		}
	}

	var _longestZigZag func(node *TreeNode)
	_longestZigZag = func(node *TreeNode) {
		branchs := nodeZzCache[node]
		if node.Left != nil {
			_longestZigZag(node.Left)
			length := 1 + nodeZzCache[node.Left][right]
			branchs[left] = length
			updateMax(length)
		} else {
			branchs[left] = 0
		}
		if node.Right != nil {
			_longestZigZag(node.Right)
			length := 1 + nodeZzCache[node.Right][left]
			branchs[right] = length
			updateMax(length)
		} else {
			branchs[right] = 0
		}
		nodeZzCache[node] = branchs
	}
	_longestZigZag(root)

	return max
}
