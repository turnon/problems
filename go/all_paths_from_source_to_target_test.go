package main

import "testing"

func TestAllPathsSourceTarget1(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}})
	t.Log(paths)
}

func TestAllPathsSourceTarget2(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})
	t.Log(paths)
}

func TestAllPathsSourceTarget3(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{1}, {}})
	t.Log(paths)
}

func TestAllPathsSourceTarget4(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{1, 2, 3}, {2}, {3}, {}})
	t.Log(paths)
}

func TestAllPathsSourceTarget5(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{1, 3}, {2}, {3}, {}})
	t.Log(paths)
}

func TestAllPathsSourceTarget6(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {}, {4}, {}})
	t.Log(paths)
}

func TestAllPathsSourceTarget7(t *testing.T) {
	paths := allPathsSourceTarget([][]int{{2}, {}, {1}})
	t.Log(paths)
}

func TestAllPathsSourceTarget99(t *testing.T) {
	a := make([]int, 0, 2)
	b := append(a, 7)
	c := append(a, 8)
	t.Log(b, c)
}
