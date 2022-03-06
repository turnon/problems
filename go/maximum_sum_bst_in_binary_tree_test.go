package main

import "testing"

func TestMaxSumBST1(t *testing.T) {
	tree := TreeNode{
		1,
		&TreeNode{
			4,
			&TreeNode{2, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{
			3,
			&TreeNode{2, nil, nil},
			&TreeNode{
				5,
				&TreeNode{4, nil, nil},
				&TreeNode{6, nil, nil}},
		},
	}

	sum := maxSumBST(&tree)
	t.Log(sum)
}

func TestMaxSumBST2(t *testing.T) {
	tree := TreeNode{
		1,
		nil,
		&TreeNode{
			10,
			&TreeNode{-5, nil, nil},
			&TreeNode{20, nil, nil},
		},
	}

	sum := maxSumBST(&tree)
	t.Log(sum)
}
