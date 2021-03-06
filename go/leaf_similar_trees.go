package main

/**
LC#872 https://leetcode.com/problems/leaf-similar-trees/
判断两树的叶子序列是否一致
思路：按相同方式遍历两树树，遇到叶子结点则收集到slice中。判断两树的slice是否一致
**/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ls1 := leaves(root1)
	ls2 := leaves(root2)
	return sameLeaves(ls1, ls2)
}

func leaves(root *TreeNode) []int {
	var inOrder func(node *TreeNode)

	ls := []int{}
	inOrder = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			ls = append(ls, node.Val)
			return
		}
		if node.Left != nil {
			inOrder(node.Left)
		}
		if node.Right != nil {
			inOrder(node.Right)
		}
	}
	inOrder(root)
	return ls
}

func sameLeaves(ls1 []int, ls2 []int) bool {
	if len(ls1) != len(ls2) {
		return false
	}
	for idx, v := range ls1 {
		if v != ls2[idx] {
			return false
		}
	}
	return true
}
