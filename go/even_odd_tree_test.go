package main

import "testing"

func TestIsEvenOddTree1(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			10,
			&TreeNode{
				3,
				&TreeNode{
					12,
					nil,
					nil,
				},
				&TreeNode{
					8,
					nil,
					nil,
				},
			},
			nil,
		},
		&TreeNode{
			4,
			&TreeNode{
				7,
				&TreeNode{
					6,
					nil,
					nil,
				},
				nil,
			},
			&TreeNode{
				9,
				nil,
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
		},
	}

	if !isEvenOddTree(tree) {
		t.Error("oh no !")
	}
}

func TestIsEvenOddTree2(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			&TreeNode{
				7,
				nil,
				nil,
			},
			nil,
		},
	}

	if isEvenOddTree(tree) {
		t.Error("oh no !")
	}
}

func TestIsEvenOddTree3(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			9,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			1,
			&TreeNode{
				7,
				nil,
				nil,
			},
			nil,
		},
	}

	if isEvenOddTree(tree) {
		t.Error("oh no !")
	}
}

func TestIsEvenOddTree4(t *testing.T) {
	tree := &TreeNode{
		1,
		nil,
		nil,
	}

	if !isEvenOddTree(tree) {
		t.Error("oh no !")
	}
}

func TestIsEvenOddTree5(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			8,
			&TreeNode{
				1,
				&TreeNode{
					30,
					&TreeNode{
						17,
						nil,
						nil,
					},
					nil,
				},
				&TreeNode{
					20,
					nil,
					nil,
				},
			},
			&TreeNode{
				3,
				&TreeNode{
					18,
					nil,
					nil,
				},
				&TreeNode{
					16,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			6,
			&TreeNode{
				9,
				&TreeNode{
					12,
					nil,
					nil,
				},
				&TreeNode{
					10,
					nil,
					nil,
				},
			},
			&TreeNode{
				11,
				&TreeNode{
					4,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
		},
	}

	if !isEvenOddTree(tree) {
		t.Error("oh no !")
	}
}

func TestIsEvenOddTree6(t *testing.T) {
	tree := &TreeNode{
		2,
		&TreeNode{
			12,
			&TreeNode{
				5,
				&TreeNode{
					18,
					nil,
					nil,
				},
				&TreeNode{
					16,
					nil,
					nil,
				},
			},
			&TreeNode{
				9,
				nil,
				nil,
			},
		},
		&TreeNode{
			8,
			nil,
			nil,
		},
	}

	if isEvenOddTree(tree) {
		t.Error("oh no !")
	}
}
