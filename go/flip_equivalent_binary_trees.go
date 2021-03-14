package main

/**
LC#951 https://leetcode.com/problems/flip-equivalent-binary-trees/
翻转相等
思路：先序遍历，判断当前结点是否相等，再判断子结点左左右右、左右右左是否相等
**/
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	if (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)) {
		return true
	}
	return false
}
