package main

import "testing"

func TestMaxLevelSum1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{7,
			&TreeNode{7,
				nil,
				nil},
			&TreeNode{-8,
				nil,
				nil}},
		&TreeNode{0,
			nil,
			nil},
	}
	if l := maxLevelSum(&p); l != 2 {
		t.Error(l)
	}
}

func TestMaxLevelSum2(t *testing.T) {
	p := TreeNode{989,
		nil,
		&TreeNode{10250,
			nil,
			nil},
	}
	if l := maxLevelSum(&p); l != 2 {
		t.Error(l)
	}
}
