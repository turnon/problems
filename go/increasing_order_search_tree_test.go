package main

import "testing"

func TestIncreasingBST(t *testing.T) {
	p := TreeNode{5,
		&TreeNode{3,
			&TreeNode{2,
				&TreeNode{1, nil, nil},
				nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{6,
			nil,
			&TreeNode{8,
				&TreeNode{7, nil, nil},
				&TreeNode{9, nil, nil}}},
	}
	preOrderP(&p)
	q := increasingBST(&p)
	preOrderP(q)
}
func TestIncreasingBST1(t *testing.T) {
	p := TreeNode{5,
		&TreeNode{1,
			nil, nil},
		&TreeNode{7,
			nil, nil},
	}
	preOrderP(&p)
	q := increasingBST(&p)
	preOrderP(q)
}
