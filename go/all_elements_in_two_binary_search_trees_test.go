package main

import (
	"fmt"
	"testing"
)

func TestGetAllElements1(t *testing.T) {
	p1 := TreeNode{
		2,
		&TreeNode{
			1,
			nil, nil,
		},
		&TreeNode{4,
			nil, nil,
		},
	}
	p2 := TreeNode{
		1,
		&TreeNode{
			0,
			nil, nil,
		},
		&TreeNode{3,
			nil, nil,
		},
	}
	list := getAllElements(&p1, &p2)
	fmt.Println(list)
}

func TestGetAllElements2(t *testing.T) {
	p1 := TreeNode{
		0,
		&TreeNode{
			-10,
			nil, nil,
		},
		&TreeNode{10,
			nil, nil,
		},
	}
	p2 := TreeNode{
		5,
		&TreeNode{
			1,
			&TreeNode{
				0,
				nil, nil,
			}, &TreeNode{
				2,
				nil, nil,
			},
		},
		&TreeNode{7,
			nil, nil,
		},
	}
	list := getAllElements(&p1, &p2)
	fmt.Println(list)
}

func TestGetAllElements5(t *testing.T) {
	p1 := TreeNode{
		1,
		nil,
		&TreeNode{8,
			nil, nil,
		},
	}
	p2 := TreeNode{
		8,
		&TreeNode{
			1,
			nil, nil,
		}, nil,
	}
	list := getAllElements(&p1, &p2)
	fmt.Println(list)
}
