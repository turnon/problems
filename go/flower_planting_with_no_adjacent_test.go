package main

import "testing"

func TestGardenNoAdj1(t *testing.T) {
	res := gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}})
	t.Log(res)
}

func TestGardenNoAdj2(t *testing.T) {
	res := gardenNoAdj(4, [][]int{{1, 2}, {3, 4}})
	t.Log(res)
}

func TestGardenNoAdj3(t *testing.T) {
	res := gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}})
	t.Log(res)
}
