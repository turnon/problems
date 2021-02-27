package main

import "fmt"

// TreeNode is just tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderP(n *TreeNode) {
	if n == nil {
		return
	}
	arr := []int{}
	preOrderCollect(n, &arr)
	fmt.Println(arr)
}

func preOrderCollect(n *TreeNode, arr *[]int) {
	if n == nil {
		return
	}
	*arr = append(*arr, n.Val)
	preOrderCollect(n.Left, arr)
	preOrderCollect(n.Right, arr)
}
