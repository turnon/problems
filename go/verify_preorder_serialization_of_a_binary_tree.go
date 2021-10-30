package main

import (
	"strings"
)

/*
LC#331 验证先序遍历序列的正确性
https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree
思路：创建栈，压入序列的第一个结点作为根结点。
遍历序列中剩余的结点，将其赋值为栈顶结点的子结点。
如果栈顶结点的左右子结点都已赋值，则出栈（后续结点不是其子结点）。
如果当前结点不是空结点，则压栈（后续结点是其子结点）。
遍历完成后检查栈是否为空（如果非空则表示序列不完整或分支节点为空结点，序列不正确）
*/
func isValidSerialization(preorder string) bool {
	const nilNode = "#"
	if preorder == nilNode {
		return true
	}

	type btNode struct{ left, right *btNode }

	nodeVals := strings.Split(preorder, ",")
	root, nodeVals := nodeVals[0], nodeVals[1:]
	if root == nilNode {
		return false
	}
	stack := make([]*btNode, 0, len(nodeVals))
	stack = append(stack, &btNode{})

	for _, val := range nodeVals {
		length := len(stack)
		if length == 0 {
			return false
		}

		lastNode := stack[length-1]
		currNode := &btNode{}

		if lastNode.left == nil {
			lastNode.left = currNode
		} else if lastNode.right == nil {
			lastNode.right = currNode
			stack = stack[:length-1]
		}

		if val != nilNode {
			stack = append(stack, currNode)
		}
	}

	return len(stack) == 0
}
