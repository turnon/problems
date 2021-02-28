package main

import "testing"

func TestAddOneRow1(t *testing.T) {
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

	neo := addOneRow(&old, 1, 2)

	preOrderP(neo)
}

func TestAddOneRow2(t *testing.T) {
	old := TreeNode{
		4,
		&TreeNode{
			2,
			&TreeNode{3, nil, nil},
			&TreeNode{1, nil, nil},
		},
		nil,
	}

	neo := addOneRow(&old, 1, 3)

	preOrderP(neo)
}

func TestAddOneRow3(t *testing.T) {
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

	neo := addOneRow(&old, 1, 0)

	preOrderP(neo)
}
