package main

/**
LC#623 https://leetcode.com/problems/add-one-row-to-tree/submissions/
往二叉树中增加一行
思路：递归下探至指定深度增加结点即可
**/
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return nil
	}
	if d > 2 {
		d = d - 1
		addOneRow(root.Left, v, d)
		addOneRow(root.Right, v, d)
	} else if d == 2 {
		root.Left = &TreeNode{v, root.Left, nil}
		root.Right = &TreeNode{v, nil, root.Right}
	} else {
		return &TreeNode{v, root, nil}
	}
	return root
}
