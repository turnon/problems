package main

import "testing"

func TestInOrderP(t *testing.T) {
	old := TreeNode{
		4,
		&TreeNode{
			2,
			&TreeNode{3, nil, nil},
			&TreeNode{1, nil, nil},
		},
		&TreeNode{6,
			&TreeNode{5, nil, nil},
			nil,
		},
	}

	inOrderP(&old)
}
