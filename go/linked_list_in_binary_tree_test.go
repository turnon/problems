package main

import "testing"

func TestIsSubPath1(t *testing.T) {
	tree := TreeNode{
		1,
		&TreeNode{
			4,
			nil,
			&TreeNode{
				2,
				&TreeNode{
					1,
					nil, nil,
				},
				nil,
			},
		},
		&TreeNode{
			4,
			&TreeNode{
				2,
				&TreeNode{
					6,
					nil, nil,
				},
				&TreeNode{
					8,
					&TreeNode{
						1,
						nil,
						nil,
					},
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			nil,
		},
	}

	link := ListNode{4, &ListNode{2, &ListNode{8, nil}}}

	if !isSubPath(&link, &tree) {
		t.Error("no !")
	}
}

func TestIsSubPath2(t *testing.T) {
	tree := TreeNode{
		1,
		&TreeNode{
			4,
			nil,
			&TreeNode{
				2,
				&TreeNode{
					1,
					nil, nil,
				},
				nil,
			},
		},
		&TreeNode{
			4,
			&TreeNode{
				2,
				&TreeNode{
					6,
					nil, nil,
				},
				&TreeNode{
					8,
					&TreeNode{
						1,
						nil,
						nil,
					},
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			nil,
		},
	}

	link := ListNode{1, &ListNode{4, &ListNode{2, &ListNode{6, nil}}}}

	if !isSubPath(&link, &tree) {
		t.Error("no !")
	}
}

func TestIsSubPath3(t *testing.T) {
	tree := TreeNode{
		1,
		&TreeNode{
			4,
			nil,
			&TreeNode{
				2,
				&TreeNode{
					1,
					nil, nil,
				},
				nil,
			},
		},
		&TreeNode{
			4,
			&TreeNode{
				2,
				&TreeNode{
					6,
					nil, nil,
				},
				&TreeNode{
					8,
					&TreeNode{
						1,
						nil,
						nil,
					},
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			nil,
		},
	}

	link := ListNode{1, &ListNode{4, &ListNode{2, &ListNode{6, &ListNode{8, nil}}}}}

	if isSubPath(&link, &tree) {
		t.Error("no !")
	}
}
