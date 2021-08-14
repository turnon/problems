package main

import "testing"

func TestFindCenter1(t *testing.T) {
	expected := 2
	if center := findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}); center != expected {
		t.Error(center, "should be", expected)
	}
}

func TestFindCenter2(t *testing.T) {
	expected := 1
	if center := findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}); center != expected {
		t.Error(center, "should be", expected)
	}
}

func TestFindCenter3(t *testing.T) {
	expected := 18
	if center := findCenter([][]int{{1, 18}, {18, 2}, {3, 18}, {18, 4}, {18, 5}, {6, 18}, {18, 7}, {18, 8}, {18, 9}, {18, 10}, {18, 11}, {12, 18}, {18, 13}, {18, 14}, {15, 18}, {16, 18}, {17, 18}}); center != expected {
		t.Error(center, "should be", expected)
	}
}

func TestFindCenter4(t *testing.T) {
	expected := 16
	if center := findCenter([][]int{{1, 16}, {16, 2}, {3, 16}, {4, 16}, {16, 5}, {16, 6}, {7, 16}, {16, 8}, {9, 16}, {10, 16}, {16, 11}, {12, 16}, {16, 13}, {16, 14}, {15, 16}, {16, 17}}); center != expected {
		t.Error(center, "should be", expected)
	}
}

func TestFindCenter5(t *testing.T) {
	expected := 9
	if center := findCenter([][]int{{9, 1}, {9, 2}, {9, 3}, {4, 9}, {9, 5}, {6, 9}, {9, 7}, {9, 8}, {10, 9}, {11, 9}, {12, 9}}); center != expected {
		t.Error(center, "should be", expected)
	}
}

func TestFindCenter6(t *testing.T) {
	expected := 3
	if center := findCenter([][]int{{1, 3}, {2, 3}}); center != expected {
		t.Error(center, "should be", expected)
	}
}
