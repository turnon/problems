package main

/**
LC#1367 https://leetcode.com/problems/linked-list-in-binary-tree/
链表是否存在于二叉树中
思路：遍历二叉树每个结点，检查其是否等于链表的头结点，并递归检查其左右子结点是否等于链表的下一结点，如果某个树结点底下存在指定链表，则结束遍历
**/
func isSubPath(head *ListNode, root *TreeNode) bool {
	result := false

	var _isSubPath func(listNode *ListNode, treeNode *TreeNode) bool
	_isSubPath = func(listNode *ListNode, treeNode *TreeNode) bool {
		if treeNode == nil {
			return listNode == nil
		}

		if listNode == nil {
			return true
		}

		if treeNode.Val == listNode.Val && (_isSubPath(listNode.Next, treeNode.Left) || _isSubPath(listNode.Next, treeNode.Right)) {
			return true
		}

		return false
	}

	var recursive func(treeNode *TreeNode)
	recursive = func(treeNode *TreeNode) {
		if result {
			return
		}
		result = _isSubPath(head, treeNode)
		if treeNode == nil {
			return
		}
		recursive(treeNode.Left)
		recursive(treeNode.Right)

	}
	recursive(root)

	return result
}
