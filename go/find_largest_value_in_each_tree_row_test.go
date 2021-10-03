package main

import "testing"

func TestLargestValues1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{3,
			&TreeNode{5, nil, nil},
			&TreeNode{3, nil, nil}},
		&TreeNode{2,
			nil,
			&TreeNode{9, nil, nil}},
	}
	t.Log(largestValues(&p))
}

func TestLargestValues2(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	t.Log(largestValues(&p))
}

func TestLargestValues3(t *testing.T) {
	p := TreeNode{1, nil, nil}
	t.Log(largestValues(&p))
}

func TestLargestValues4(t *testing.T) {
	p := TreeNode{1, nil, &TreeNode{2, nil, nil}}
	t.Log(largestValues(&p))
}

func TestLargestValues45(t *testing.T) {

	t.Log(largestValues(nil))
}
