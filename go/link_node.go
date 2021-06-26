package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func _printListNodes(node *ListNode) {
	list := []int{}
	for ; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	fmt.Println(list)
}
