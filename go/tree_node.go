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

func inOrderP(n *TreeNode) {
	if n == nil {
		return
	}

	values := []int{}
	queue := []*TreeNode{n}
	for {
		if len(queue) == 0 {
			break
		}

		curr := queue[0]
		queue = queue[1:]
		values = append(values, curr.Val)

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	fmt.Println(values)
}
