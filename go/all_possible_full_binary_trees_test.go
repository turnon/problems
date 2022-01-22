package main

import (
	"testing"
)

func TestAllPossibleFBT1(t *testing.T) {
	result := allPossibleFBT(7)
	t.Log(result)
}

func TestAllPossibleFBT2(t *testing.T) {
	result := allPossibleFBT(2)
	t.Log(result)
}
