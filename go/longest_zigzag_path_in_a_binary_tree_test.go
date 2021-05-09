package main

import "testing"

func TestLongestZigZag1(t *testing.T) {
	p := TreeNode{
		1,
		nil,
		&TreeNode{
			1,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				1,
				&TreeNode{
					1,
					nil,
					&TreeNode{
						1,
						nil,
						&TreeNode{
							1,
							nil,
							nil,
						},
					},
				},
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
	}

	l := longestZigZag(&p)
	if l != 3 {
		t.Error(l)
	}
}

func TestLongestZigZag2(t *testing.T) {
	p := TreeNode{
		1,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				1,
				&TreeNode{
					1,
					nil,
					&TreeNode{
						1,
						nil,
						nil,
					},
				},
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}

	l := longestZigZag(&p)
	if l != 4 {
		t.Error(l)
	}
}

func TestLongestZigZag3(t *testing.T) {
	p := TreeNode{
		1,
		nil,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				1,
				nil,
				&TreeNode{
					1,
					nil,
					&TreeNode{
						1,
						nil,
						&TreeNode{
							1,
							&TreeNode{
								1,
								nil,
								nil,
							},
							nil,
						},
					},
				},
			},
		},
	}

	l := longestZigZag(&p)
	if l != 2 {
		t.Error(l)
	}
}
