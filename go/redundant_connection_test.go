package main

import "testing"

func TestFindRedundantConnection1(t *testing.T) {
	res := findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})
	cmpFindRedundantConnection(t, res, []int{2, 3})
}

func TestFindRedundantConnection2(t *testing.T) {
	res := findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})
	cmpFindRedundantConnection(t, res, []int{1, 4})
}

func TestFindRedundantConnection3(t *testing.T) {
	res := findRedundantConnection([][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}})
	cmpFindRedundantConnection(t, res, []int{2, 5})
}

func cmpFindRedundantConnection(t *testing.T, a, b []int) {
	if a == nil {
		t.Error("nil should be", b)
	}
	if a[0] != b[0] || a[1] != b[1] {
		t.Error(a, "should be", b)
	}
}
