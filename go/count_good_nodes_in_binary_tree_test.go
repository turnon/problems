package main

import "testing"

func TestGoodNodes1(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{
			1,
			&TreeNode{
				3,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			4,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
	}

	if count := goodNodes(tree); count != 4 {
		t.Error(count)
	}
}

func TestGoodNodes2(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{
			3,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		nil,
	}

	if count := goodNodes(tree); count != 3 {
		t.Error(count)
	}
}

func TestGoodNodes3(t *testing.T) {
	tree := &TreeNode{
		1,
		nil,
		nil,
	}

	if count := goodNodes(tree); count != 1 {
		t.Error(count)
	}
}
