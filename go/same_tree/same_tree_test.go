package sametree

import "testing"

func TestSameTree1(t *testing.T) {
	p := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	if !isSameTree(&p, &q) {
		t.Error("oh no")
	}
}

func TestSameTree2(t *testing.T) {
	p := TreeNode{1, &TreeNode{2, nil, nil}, nil}
	q := TreeNode{1, nil, &TreeNode{2, nil, nil}}
	if isSameTree(&p, &q) {
		t.Error("oh no")
	}
}
