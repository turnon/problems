package main

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	tree := TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{4,
				&TreeNode{7, nil, nil},
				nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3,
			&TreeNode{6,
				nil,
				&TreeNode{8, nil, nil}},
			nil,
		},
	}

	if sum := deepestLeavesSum(&tree); sum != 15 {
		t.Error(sum)
	}
}
