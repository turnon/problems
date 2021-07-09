package main

import "testing"

func TestFindMinHeightTrees1(t *testing.T) {
	result := findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	t.Log(result)
}

func TestFindMinHeightTrees2(t *testing.T) {
	result := findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
	t.Log(result)
}

func TestFindMinHeightTrees3(t *testing.T) {
	result := findMinHeightTrees(1, [][]int{})
	t.Log(result)
}

func TestFindMinHeightTrees4(t *testing.T) {
	result := findMinHeightTrees(2, [][]int{{0, 1}})
	t.Log(result)
}
