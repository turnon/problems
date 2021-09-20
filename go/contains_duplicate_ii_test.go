package main

import "testing"

func TestContainsNearbyDuplicate1(t *testing.T) {
	if !containsNearbyDuplicate([]int{1, 2, 3, 1}, 3) {
		t.Error("should be ", true)
	}
}

func TestContainsNearbyDuplicate2(t *testing.T) {
	if !containsNearbyDuplicate([]int{1, 0, 1, 1}, 1) {
		t.Error("should be ", true)
	}
}

func TestContainsNearbyDuplicate3(t *testing.T) {
	if containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2) {
		t.Error("should be ", false)
	}
}

func TestContainsNearbyDuplicate4(t *testing.T) {
	if containsNearbyDuplicate([]int{1}, 1) {
		t.Error("should be ", false)
	}
}

func TestContainsNearbyDuplicate5(t *testing.T) {
	if containsNearbyDuplicate([]int{1, 2}, 2) {
		t.Error("should be ", false)
	}
}
