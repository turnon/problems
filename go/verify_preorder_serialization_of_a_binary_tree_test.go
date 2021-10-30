package main

import "testing"

func TestIsValidSerialization1(t *testing.T) {
	if !isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#") {
		t.Error("should be ", true)
	}
}

func TestIsValidSerialization2(t *testing.T) {
	if isValidSerialization("1,#") {
		t.Error("should be ", false)
	}
}

func TestIsValidSerialization3(t *testing.T) {
	if isValidSerialization("9,#,#,1") {
		t.Error("should be ", false)
	}
}
