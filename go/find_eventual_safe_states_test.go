package main

import "testing"

func TestEventualSafeNodes1(t *testing.T) {
	result := eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}})
	testEventualSafeNodes(t, result, []int{2, 4, 5, 6})

}
func TestEventualSafeNodes2(t *testing.T) {
	result := eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}})
	testEventualSafeNodes(t, result, []int{4})
}

func testEventualSafeNodes(t *testing.T, actual, expected []int) {
	eq := true
	if len(expected) != len(actual) {
		eq = false
	}
	if eq {
		for i, v := range expected {
			if v != actual[i] {
				eq = false
			}
		}
	}
	if !eq {
		t.Log(len(actual), "->", actual)
	}
}
