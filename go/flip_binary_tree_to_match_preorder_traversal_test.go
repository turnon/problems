package main

import (
	"sort"
	"testing"
)

func TestFlipMatchVoyage(t *testing.T) {
	sameArray := func(actual, expected []int) bool {
		if len(actual) != len(expected) {
			return false
		}

		sort.Ints(actual)
		sort.Ints(expected)

		for i, a := range actual {
			if expected[i] != a {
				return false
			}
		}

		return true
	}

	testcases := []struct {
		root     *TreeNode
		preorder []int
		expected []int
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, nil}, []int{2, 1}, []int{-1}},
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, []int{1, 3, 2}, []int{1}},
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, []int{1, 2, 3}, []int{}},
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}, []int{1, 3, 4, 5, 2}, []int{1}},
	}

	for i, testcase := range testcases {
		actual := flipMatchVoyage(testcase.root, testcase.preorder)

		if !sameArray(actual, testcase.expected) {
			t.Errorf("testcase %v failed : expecting %v, but actually %v", i+1, testcase.expected, actual)
		}
	}
}
