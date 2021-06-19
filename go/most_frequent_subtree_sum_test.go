package main

import (
	"sort"
	"testing"
)

func TestFindFrequentTreeSum1(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			-3,
			nil,
			nil,
		},
	}

	res := findFrequentTreeSum(tree)
	sort.Ints(res)
	if res[0] != -3 || res[1] != 2 || res[2] != 4 {
		t.Error(res)
	}
}

func TestFindFrequentTreeSum2(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			-5,
			nil,
			nil,
		},
	}

	res := findFrequentTreeSum(tree)
	sort.Ints(res)
	if res[0] != 2 {
		t.Error(res)
	}
}
