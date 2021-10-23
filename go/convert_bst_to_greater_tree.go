package main

/*
LC#538 将平衡二叉树的每个结点值变成“原值加上更大的那些结点的值之和”
https://leetcode.com/problems/convert-bst-to-greater-tree/
思路：按右结点、根结点、左结点的顺序递归遍历，
在处理根节点时，将其值改为“该结点值加上前一个处理到的结点的值”，并标记为“前一个处理到的结点”
*/
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var lastNode *TreeNode
	var sortInOrder func(node *TreeNode)
	sortInOrder = func(node *TreeNode) {
		if node.Right != nil {
			sortInOrder(node.Right)
		}

		if lastNode != nil {
			node.Val += lastNode.Val
		}
		lastNode = node
		if node.Left != nil {
			sortInOrder(node.Left)
		}
	}
	sortInOrder(root)

	return root
}
