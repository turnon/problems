package main

import "testing"

func TestFindJudge1(t *testing.T) {
	const exp = 2
	res := findJudge(2, [][]int{{1, 2}})
	if res != exp {
		t.Error(res, "should be", exp)
	}
}

func TestFindJudge2(t *testing.T) {
	const exp = 3
	res := findJudge(3, [][]int{{1, 3}, {2, 3}})
	if res != exp {
		t.Error(res, "should be", exp)
	}
}

func TestFindJudge3(t *testing.T) {
	const exp = -1
	res := findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})
	if res != exp {
		t.Error(res, "should be", exp)
	}
}

func TestFindJudge4(t *testing.T) {
	const exp = -1
	res := findJudge(3, [][]int{{1, 2}, {2, 3}})
	if res != exp {
		t.Error(res, "should be", exp)
	}
}

func TestFindJudge5(t *testing.T) {
	const exp = 3
	res := findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}})
	if res != exp {
		t.Error(res, "should be", exp)
	}
}

func TestFindJudge6(t *testing.T) {
	const exp = 1
	res := findJudge(1, [][]int{})
	if res != exp {
		t.Error(res, "should be", exp)
	}
}
