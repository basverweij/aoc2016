package main

import (
	"testing"
)

func TestCalcPassword(t *testing.T) {
	pwd := calcPassword("abc")
	if pwd != "05ace8e3" {
		t.Errorf("expected password '05ace8e3' for input 'abc', but got '%s'", pwd)
	}
}

func testCalcHash(t *testing.T, input string, index int, expectedHash string) {
	actualHash := calcHash(input, index)
	if actualHash != expectedHash {
		t.Errorf(
			"expected hash '%s' for input '%s' and index %d, but got '%s'",
			expectedHash, input, index, actualHash)
	}
}

func TestCalcHash(t *testing.T) {
	testCalcHash(t, "abc", 3231929, "0000015")
	testCalcHash(t, "abc", 5017308, "000008f")
	testCalcHash(t, "abc", 5278568, "00000f9")
	testCalcHash(t, "abc", 5357525, "000004e")
}
