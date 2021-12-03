package main

/*
LC#701 往二叉查找树中插入结点
https://leetcode.com/problems/insert-into-a-binary-search-tree/
思路：下探至空缺的左结点指针或右结点指针，赋值即可
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{val, nil, nil}
	if root == nil {
		return node
	}
	ptr := &root
	for {
		curr := *ptr
		if curr == nil {
			break
		}
		if val < curr.Val {
			ptr = &curr.Left
		} else {
			ptr = &curr.Right
		}
	}
	*ptr = node
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{val, nil, nil}
	if root == nil {
		return node
	}
	curr := root
	for {
		if val < curr.Val {
			if curr.Left == nil {
				curr.Left = node
				break
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = node
				break
			}
			curr = curr.Right
		}
	}
	return root
}
