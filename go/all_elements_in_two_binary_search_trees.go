package main

/**
LC#1305 https://leetcode.com/problems/all-elements-in-two-binary-search-trees
二叉查找树归并排序成数组
思路：中序遍历得两切片，再归并排序
**/
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var recursive func(node *TreeNode, list []int) []int
	recursive = func(node *TreeNode, list []int) []int {
		if node == nil {
			return list
		}
		list = recursive(node.Left, list)
		list = append(list, node.Val)
		list = recursive(node.Right, list)
		return list
	}

	list1 := recursive(root1, []int{})
	list2 := recursive(root2, []int{})

	list := make([]int, 0, len(list1)+len(list2))
	l1Idx := 0
	l1len := len(list1)
	l2Idx := 0
	l2len := len(list2)
	for {
		if l1Idx >= l1len || l2Idx >= l2len {
			break
		}
		if list1[l1Idx] > list2[l2Idx] {
			list = append(list, list2[l2Idx])
			l2Idx++
		} else {
			list = append(list, list1[l1Idx])
			l1Idx++
		}
	}
	if l1Idx < l1len {
		list = append(list, list1[l1Idx:]...)
	} else {
		list = append(list, list2[l2Idx:]...)
	}

	return list
}
