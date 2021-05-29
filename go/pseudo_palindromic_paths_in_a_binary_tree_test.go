package main

import "testing"

func TestPseudoPalindromicPaths1(t *testing.T) {
	tree := &TreeNode{
		2,
		&TreeNode{
			3,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				1,
				nil,
				nil,
			},
		},
		&TreeNode{
			1,
			nil,
			&TreeNode{
				1,
				nil,
				nil,
			},
		},
	}

	if c := pseudoPalindromicPaths(tree); c != 2 {
		t.Error(c)
	}
}

func TestPseudoPalindromicPaths2(t *testing.T) {
	tree := &TreeNode{
		2,
		&TreeNode{
			1,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}

	if c := pseudoPalindromicPaths(tree); c != 1 {
		t.Error(c)
	}
}

func TestPseudoPalindromicPaths3(t *testing.T) {
	tree := &TreeNode{
		9,
		nil,
		nil,
	}

	if c := pseudoPalindromicPaths(tree); c != 1 {
		t.Error(c)
	}
}
