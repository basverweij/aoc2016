package main

import (
	"testing"
)

func TestCalcPassword(t *testing.T) {
	pwd := calcPassword("abc")
	if pwd != "18f47a30" {
		t.Errorf("expected password '18f47a30' for input 'abc', but got '%s'", pwd)
	}
}

func testCalcHash(t *testing.T, input string, index int, expectedHash string) {

}

func TestCalcHash(t *testing.T) {
	testCalcHash(t, "abc", 3231929, "000001")
	testCalcHash(t, "abc", 5017308, "000008")
	testCalcHash(t, "abc", 5278568, "00000f")
}
