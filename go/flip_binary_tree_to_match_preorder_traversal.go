package main

/*
LC#971 https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal
翻转二叉树以适应先序遍历序列
思路：先左后右递归地先序遍历二叉树的过程中，推进序列指针，比较当前结点值和指针值是否相等，
若不等，则回退指针，改为先右后左再次递归先序遍历，此时记下当前被翻转的结点，
遍历完成后，还要检查序列是否已到尽头，若是，则返回记下的结点集合，否则返回-1
*/
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	currentIdx := -1
	voyageLen := len(voyage)
	result := []int{}
	getCurrVal := func() int {
		if voyageLen > currentIdx {
			return voyage[currentIdx]
		}
		return 0
	}

	var walk func(node *TreeNode) bool
	walk = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		currentIdx += 1
		currVal := getCurrVal()
		if node.Val != currVal {
			currentIdx -= 1
			return false
		}

		if walk(node.Left) && walk(node.Right) {
			return true
		}

		result = append(result, node.Val)
		if walk(node.Right) && walk(node.Left) {
			return true
		}

		return false
	}

	if walk(root) && currentIdx+1 == voyageLen {
		return result
	}

	return []int{-1}
}
