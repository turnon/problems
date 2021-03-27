package main

import "testing"

func TestBstFromPreorder1(t *testing.T) {
	args := []int{8, 5, 1, 7, 10, 12}
	n := bstFromPreorder(args)
	preOrderP(n)
}

func TestBstFromPreorder2(t *testing.T) {
	args := []int{1, 3}
	n := bstFromPreorder(args)
	preOrderP(n)
	preOrderP(n.Right)
}

func TestBstFromPreorder3(t *testing.T) {
	args := []int{4, 2}
	n := bstFromPreorder(args)
	preOrderP(n)
	preOrderP(n.Left)
}
