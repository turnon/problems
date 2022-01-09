package main

/*
LC#965 单值树
https://leetcode.com/problems/univalued-binary-tree
思路：先序遍历，如果结点的值不等于根的值，马上返回false
*/
func isUnivalTree(root *TreeNode) bool {
	v := root.Val
	var walk func(node *TreeNode) bool
	walk = func(node *TreeNode) bool {
		if node.Val != v {
			return false
		}
		if node.Left != nil && !walk(node.Left) {
			return false
		}
		if node.Right != nil && !walk(node.Right) {
			return false
		}
		return true
	}
	return walk(root)
}
