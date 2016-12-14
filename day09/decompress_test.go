package main

import "testing"

func assertEquals(t *testing.T, expected, actual interface{}, msg string) {
	if expected == actual {
		return
	}

	t.Errorf("%s: expected '%v', but got '%v'", msg, expected, actual)
}

func assertNotEquals(t *testing.T, notExpected, actual interface{}, msg string) {
	if notExpected != actual {
		return
	}

	t.Errorf("%s: did not expect '%v'", msg, notExpected)
}

func TestDecompressEmptyString(t *testing.T) {
	c, err := decompress("")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 0, c, "output string")
}

func TestDecompressStringWithoutMarkers(t *testing.T) {
	c, err := decompress("ADVENT")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 6, c, "output string")
}

func TestDecompressStringWithSingleMarker(t *testing.T) {
	c, err := decompress("A(1x5)BC")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 7, c, "output string")
}

func TestDecompressStringWithMarkerAtStart(t *testing.T) {
	c, err := decompress("(3x3)XYZ")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 9, c, "output string")
}

func TestDecompressStringWithMultipleMarkers(t *testing.T) {
	c, err := decompress("A(2x2)BCD(2x2)EFG")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 11, c, "output string")
}

func TestDecompressStringWithRecursiveMarker(t *testing.T) {
	c, err := decompress("(6x1)(1x3)A")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 3, c, "output string")
}

func TestDecompressStringWithRepeatedRecursiveMarker(t *testing.T) {
	c, err := decompress("X(8x2)(3x3)ABCY")
	assertEquals(t, nil, err, "error")
	assertEquals(t, 20, c, "output string")
}

func TestScanIntEmptySlice(t *testing.T) {
	_, _, err := scanInt([]rune(""), 0, '#')
	assertNotEquals(t, nil, err, "error")
}

func TestScanIntWithoutStopRune(t *testing.T) {
	_, _, err := scanInt([]rune("ADVENT"), 0, '#')
	assertNotEquals(t, nil, err, "error")
}

func TestScanIntForX(t *testing.T) {
	value, end, err := scanInt([]rune("(1x2)"), 1, 'x')
	assertEquals(t, nil, err, "error")
	assertEquals(t, 1, value, "value")
	assertEquals(t, 2, end, "end")
}

func TestScanIntForClosingParenthesis(t *testing.T) {
	value, end, err := scanInt([]rune("(1x2)"), 3, ')')
	assertEquals(t, nil, err, "error")
	assertEquals(t, 2, value, "value")
	assertEquals(t, 4, end, "end")
}
