package main

import "testing"

func TestCanVisitAllRooms1(t *testing.T) {
	if !canVisitAllRooms([][]int{{1}, {2}, {3}, {}}) {
		t.Error("should be true")
	}
}

func TestCanVisitAllRooms2(t *testing.T) {
	if canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}) {
		t.Error("should be false")
	}
}
