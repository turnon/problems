package main

import "testing"

func TestFindCircleNum1(t *testing.T) {
	if count := findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}); count != 2 {
		t.Error(count)
	}
}

func TestFindCircleNum2(t *testing.T) {
	if count := findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}); count != 3 {
		t.Error(count)
	}
}

func TestFindCircleNum3(t *testing.T) {
	if count := findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}); count != 1 {
		t.Error(count)
	}
}
