package main

import "testing"

func TestWidthOfBinaryTree1(t *testing.T) {
	p := &TreeNode{2,
		&TreeNode{1, &TreeNode{3, nil, nil}, nil},
		&TreeNode{4, &TreeNode{5, nil, nil}, nil},
	}
	if r := widthOfBinaryTree(p); r != 3 {
		t.Error(r, 3)
	}
}

func TestWidthOfBinaryTree2(t *testing.T) {
	p := &TreeNode{1,
		&TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, nil}, nil},
		&TreeNode{2, nil, &TreeNode{9, nil, &TreeNode{7, nil, nil}}},
	}
	if r := widthOfBinaryTree(p); r != 8 {
		t.Error(r, 8)
	}
}
