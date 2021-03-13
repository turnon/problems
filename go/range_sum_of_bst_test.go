package main

import (
	"strconv"
	"testing"
)

func TestRangeSumBST1(t *testing.T) {
	p := TreeNode{10,
		&TreeNode{5,
			&TreeNode{3,
				nil, nil},
			&TreeNode{7,
				nil, nil}},
		&TreeNode{15,
			nil,
			&TreeNode{18,
				nil, nil},
		},
	}

	res := rangeSumBST(&p, 7, 15)
	if res != 32 {
		t.Error("7~15: " + strconv.Itoa(res))
	}
}

func TestRangeSumBST2(t *testing.T) {
	p := TreeNode{10,
		&TreeNode{5,
			&TreeNode{3,
				&TreeNode{1,
					nil, nil}, nil},
			&TreeNode{7,
				&TreeNode{6,
					nil, nil}, nil}},
		&TreeNode{15,
			&TreeNode{13,
				nil, nil},
			&TreeNode{18,
				nil, nil},
		},
	}

	res := rangeSumBST(&p, 6, 10)
	if res != 23 {
		t.Error("6-10: " + strconv.Itoa(res))
	}
}
