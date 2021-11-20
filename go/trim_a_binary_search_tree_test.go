package main

import "testing"

func TestTrimBST1(t *testing.T) {
	p := &TreeNode{1,
		&TreeNode{0, nil, nil},
		&TreeNode{2, nil, nil},
	}
	trimBST(p, 1, 2)
	inOrderP(p)
}

func TestTrimBST2(t *testing.T) {
	p := &TreeNode{3,
		&TreeNode{0, nil, &TreeNode{2, &TreeNode{1, nil, nil}, nil}},
		&TreeNode{4, nil, nil},
	}
	trimBST(p, 1, 3)
	inOrderP(p)
}
