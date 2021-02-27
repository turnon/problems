package main

import "testing"

func TestMergeTrees1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{3,
			&TreeNode{5, nil, nil}, nil},
		&TreeNode{2, nil, nil},
	}
	q := TreeNode{
		2,
		&TreeNode{1, nil, &TreeNode{4, nil, nil}},
		&TreeNode{3, nil, &TreeNode{7, nil, nil}},
	}
	r := mergeTrees(&p, &q)
	preOrderP(r)
}

func TestMergeTrees2(t *testing.T) {
	p := TreeNode{
		1,
		nil,
		nil,
	}
	q := TreeNode{
		1,
		&TreeNode{2, nil, nil},
		nil,
	}
	r := mergeTrees(&p, &q)
	preOrderP(r)
}
