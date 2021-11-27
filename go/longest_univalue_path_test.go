package main

import "testing"

func TestLongestUnivaluePath1(t *testing.T) {
	p := &TreeNode{5,
		&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{5, nil, &TreeNode{5, nil, nil}},
	}
	if r := longestUnivaluePath(p); r != 2 {
		t.Error(r, "should be 2")
	}
}

func TestLongestUnivaluePath2(t *testing.T) {
	p := &TreeNode{1,
		&TreeNode{4, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{5, nil, &TreeNode{5, nil, nil}},
	}
	if r := longestUnivaluePath(p); r != 2 {
		t.Error(r, "should be 2")
	}
}
