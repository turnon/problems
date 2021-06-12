package main

import "testing"

func TestCountPairs1(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	if c := countPairs(tree, 3); c != 1 {
		t.Error(1, c)
	}
}

func TestCountPairs2(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			&TreeNode{
				6,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	if c := countPairs(tree, 3); c != 2 {
		t.Error(2, c)
	}
}

func TestCountPairs3(t *testing.T) {
	tree := &TreeNode{
		7,
		&TreeNode{
			1,
			&TreeNode{
				6,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			4,
			&TreeNode{
				5,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
		},
	}

	if c := countPairs(tree, 3); c != 1 {
		t.Error(1, c)
	}
}

func TestCountPairs4(t *testing.T) {
	tree := &TreeNode{
		100,
		nil,
		nil,
	}

	if c := countPairs(tree, 1); c != 0 {
		t.Error(0, c)
	}
}

func TestCountPairs5(t *testing.T) {
	tree := &TreeNode{
		100,
		&TreeNode{
			100,
			nil,
			nil,
		},
		&TreeNode{
			100,
			nil,
			nil,
		},
	}

	if c := countPairs(tree, 2); c != 1 {
		t.Error(1, c)
	}
}
