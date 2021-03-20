package main

import "testing"

func TestIsCompleteTree1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				nil, nil},
			&TreeNode{5,
				nil, nil}},
		&TreeNode{3,
			&TreeNode{6,
				nil, nil},
			nil,
		},
	}
	if !isCompleteTree(&p) {
		t.Error("oh no!")
	}
}

func TestIsCompleteTree2(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				nil, nil},
			&TreeNode{5,
				nil, nil}},
		&TreeNode{3,
			nil,
			&TreeNode{7,
				nil, nil},
		},
	}
	if isCompleteTree(&p) {
		t.Error("oh no!")
	}
}

func TestIsCompleteTree3(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			&TreeNode{3,
				nil, nil},
			nil,
		},
		nil,
	}
	if isCompleteTree(&p) {
		t.Error("oh no!")
	}
}
