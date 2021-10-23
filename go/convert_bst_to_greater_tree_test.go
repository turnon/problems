package main

import "testing"

func TestConvertBST1(t *testing.T) {
	p := TreeNode{4,
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}},
		&TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, &TreeNode{8, nil, nil}}},
	}
	convertBST(&p)
	inOrderP(&p)
}
