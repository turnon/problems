package main

import "testing"

func TestPrintTree1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2, nil, nil},
		nil,
	}
	testPrintTree(printTree(&p), t)
}

func TestPrintTree2(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2, nil, &TreeNode{4, nil, nil}},
		&TreeNode{3, nil, nil},
	}
	testPrintTree(printTree(&p), t)
}

func TestPrintTree3(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2, nil, &TreeNode{4, nil, nil}},
		&TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, nil}, nil},
	}
	testPrintTree(printTree(&p), t)
}

func testPrintTree(tree [][]string, t *testing.T) {
	for _, l := range tree {
		t.Log(l)
	}
}
