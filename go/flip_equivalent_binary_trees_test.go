package main

import "testing"

func TestFlipEquiv1(t *testing.T) {
	p := TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				nil, nil},
			&TreeNode{5,
				&TreeNode{7,
					nil, nil},
				&TreeNode{8,
					nil, nil}}},
		&TreeNode{3,
			&TreeNode{6,
				nil, nil}, nil,
		},
	}

	q := TreeNode{1,
		&TreeNode{3,
			nil,
			&TreeNode{6,
				nil, nil},
		},
		&TreeNode{2,
			&TreeNode{4,
				nil, nil},
			&TreeNode{5,
				&TreeNode{8,
					nil, nil},
				&TreeNode{7,
					nil, nil}}},
	}

	if res := flipEquiv(&p, &q); !res {
		t.Error("wtf")
	}
}

func TestFlipEquiv2(t *testing.T) {
	if res := flipEquiv(nil, nil); !res {
		t.Error("wtf")
	}
}

func TestFlipEquiv3(t *testing.T) {
	if res := flipEquiv(nil, &TreeNode{1, nil, nil}); res {
		t.Error("wtf")
	}
}

func TestFlipEquiv4(t *testing.T) {
	if res := flipEquiv(nil, &TreeNode{0, nil, &TreeNode{1, nil, nil}}); res {
		t.Error("wtf")
	}
}

func TestFlipEquiv5(t *testing.T) {
	if res := flipEquiv(&TreeNode{0, &TreeNode{1, nil, nil}, nil}, &TreeNode{0, nil, &TreeNode{1, nil, nil}}); !res {
		t.Error("wtf")
	}
}
