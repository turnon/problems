package main

import "testing"

func TestRemoveLeafNodes1(t *testing.T) {
	p1 := TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				2,
				nil, nil,
			}, nil,
		},
		&TreeNode{3,
			&TreeNode{2,
				nil, nil,
			}, &TreeNode{4,
				nil, nil,
			},
		},
	}

	p2 := removeLeafNodes(&p1, 2)
	preOrderP(p2)
}

func TestRemoveLeafNodes2(t *testing.T) {
	p1 := TreeNode{
		1,
		&TreeNode{
			3,
			&TreeNode{
				3,
				nil, nil,
			}, &TreeNode{
				2,
				nil, nil,
			},
		},
		&TreeNode{3,
			nil, nil,
		},
	}

	p2 := removeLeafNodes(&p1, 3)
	preOrderP(p2)
}

func TestRemoveLeafNodes3(t *testing.T) {
	p1 := TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				2,
				&TreeNode{
					2,
					nil, nil,
				}, nil,
			}, nil,
		},
		nil,
	}

	p2 := removeLeafNodes(&p1, 2)
	preOrderP(p2)
}

func TestRemoveLeafNodes4(t *testing.T) {
	p1 := TreeNode{
		1,
		&TreeNode{
			1,
			nil, nil,
		},
		&TreeNode{
			1,
			nil, nil,
		},
	}

	p2 := removeLeafNodes(&p1, 1)
	preOrderP(p2)
}

func TestRemoveLeafNodes5(t *testing.T) {
	p1 := TreeNode{
		1,
		&TreeNode{
			2,
			nil, nil,
		},
		&TreeNode{
			3,
			nil, nil,
		},
	}

	p2 := removeLeafNodes(&p1, 1)
	preOrderP(p2)
}
