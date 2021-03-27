package main

/**
LC#1008 根据先序遍历结果重组二叉查找树
https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal
思路：取第一个元素为根结点，在剩余切片中找出大于第一元素的索引，以其切分出两切片递归处理，作为左右子树
**/
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	firstVal := preorder[0]
	root := &TreeNode{firstVal, nil, nil}

	for idx, val := range preorder {
		if val > firstVal {
			// if idx >= 1 {
			// fmt.Println(preorder[1:idx], preorder[idx:])
			root.Left = bstFromPreorder(preorder[1:idx])
			root.Right = bstFromPreorder(preorder[idx:])
			break
			// }
		}
	}

	if root.Left == nil && root.Right == nil && len(preorder) > 1 {
		root.Left = bstFromPreorder(preorder[1:])
	}

	return root
}
