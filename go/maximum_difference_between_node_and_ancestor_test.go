package main

import "testing"

func TestMaxAncestorDiff1(t *testing.T) {
	p := TreeNode{8,
		&TreeNode{3,
			&TreeNode{1,
				nil, nil},
			&TreeNode{6,
				&TreeNode{4,
					nil, nil},
				&TreeNode{7,
					nil, nil},
			}},
		&TreeNode{10,
			nil,
			&TreeNode{14,
				&TreeNode{13,
					nil, nil},
				nil},
		},
	}
	if maxAncestorDiff(&p) != 7 {
		t.Error("not 7 !")
	}
}

func TestMaxAncestorDiff2(t *testing.T) {
	p := TreeNode{1,
		nil,
		&TreeNode{2,
			nil,
			&TreeNode{0,
				&TreeNode{3,
					nil, nil},
				nil},
		},
	}
	if maxAncestorDiff(&p) != 3 {
		t.Error("not 3 !")
	}
}

func TestMaxAncestorDiff3(t *testing.T) {
	p := TreeNode{0,
		nil,
		&TreeNode{1,
			nil,
			nil},
	}
	if maxAncestorDiff(&p) != 1 {
		t.Error("not 1 !")
	}
}
