package main

/*
LC#863 求二叉树中与目标结点距离为K的结点的集合
https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
思路：先序遍历找出目标结点，并记录根结点到它的路径；
然后对路径上的每个结点，收集与他们距离为“在路径上的索引值”的子孙结点，
如果路径长于K，则要排除根部的一些结点，如果长度小于K，则排除整个根部；
另外，总是要排除路径上的结点
*/
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	path := []*TreeNode{}
	result := []int{}

	var makePath func(node *TreeNode) bool
	makePath = func(node *TreeNode) bool {
		path = append(path, node)
		if node == target {
			return true
		}
		if node.Left != nil && makePath(node.Left) {
			return true
		}
		if node.Right != nil && makePath(node.Right) {
			return true
		}
		path = path[:len(path)-1]
		return false
	}
	makePath(root)

	var findK func(node *TreeNode, except *TreeNode, k int)
	findK = func(node *TreeNode, except *TreeNode, k int) {
		if k == 0 {
			result = append(result, node.Val)
			return
		}
		if node.Left != nil && node.Left != except {
			findK(node.Left, except, k-1)
		}
		if node.Right != nil && node.Right != except {
			findK(node.Right, except, k-1)
		}
	}

	endIdx := len(path) - 1
	useless := len(path) - k - 1
	for idx, node := range path {
		if idx < useless {
			continue
		}
		var except *TreeNode
		if idx < endIdx {
			except = path[idx+1]
		} else {
			except = nil
		}
		findK(node, except, idx-useless)
	}

	return result
}
