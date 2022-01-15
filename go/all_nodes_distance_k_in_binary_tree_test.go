package main

import "testing"

func TestDistanceK0(t *testing.T) {
	target := &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}
	p := &TreeNode{3,
		target,
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}}
	result := distanceK(p, target, 2)
	t.Log(result)
}

func TestDistanceK1(t *testing.T) {
	target := &TreeNode{2, nil, nil}
	p := &TreeNode{0,
		&TreeNode{1, &TreeNode{3, nil, nil}, target},
		nil}
	result := distanceK(p, target, 1)
	t.Log(result)
}

func TestDistanceK2(t *testing.T) {
	target := &TreeNode{1, nil, nil}
	p := target
	result := distanceK(p, target, 3)
	t.Log(result)
}
