package main

import "testing"

func TestIsCousins1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				nil, nil},
			nil,
		},
		&TreeNode{3,
			nil, nil},
	}
	if isCousins(&p, 4, 3) {
		t.Error("oh no!")
	}
}

func TestIsCousins2(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{4,
				nil, nil},
		},
		&TreeNode{3,
			nil,
			&TreeNode{5,
				nil, nil},
		},
	}
	if !isCousins(&p, 5, 4) {
		t.Error("oh no!")
	}
}

func TestIsCousins3(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{4,
				nil, nil},
		},
		&TreeNode{3,
			nil, nil},
	}
	if isCousins(&p, 2, 3) {
		t.Error("oh no!")
	}
}
