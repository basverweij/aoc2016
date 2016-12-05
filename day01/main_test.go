package main

import "testing"

func testUpdateHeading(t *testing.T, currentHeading heading, dir direction, expectedHeading heading) {
	actualHeading := updateHeading(currentHeading, dir)

	if actualHeading != expectedHeading {
		t.Errorf(
			"expected %v after updating %v with %v, but got %v",
			expectedHeading, currentHeading, dir, actualHeading)
	}
}

func TestUpdateHeading(t *testing.T) {
	testUpdateHeading(t, north, right, east)
	testUpdateHeading(t, east, right, south)
	testUpdateHeading(t, south, right, west)
	testUpdateHeading(t, west, right, north)

	testUpdateHeading(t, north, left, west)
	testUpdateHeading(t, west, left, south)
	testUpdateHeading(t, south, left, east)
	testUpdateHeading(t, east, left, north)
}

func testUpdatePosition(t *testing.T, x, y int, currentHeading heading, num int, expectedX, expectedY int) {
	actualX, actualY := updatePosition(x, y, currentHeading, num)

	if actualX != expectedX || actualY != expectedY {
		t.Errorf(
			"expected (%d,%d) after updating (%d,%d) with %d heading %v, but got (%d,%d)",
			expectedX, expectedY, x, y, num, currentHeading, actualX, actualY)
	}
}

func TestUpdatePosition(t *testing.T) {
	testUpdatePosition(t, 0, 0, north, 1, 0, 1)
	testUpdatePosition(t, 0, 0, east, 1, 1, 0)
	testUpdatePosition(t, 0, 0, south, 1, 0, -1)
	testUpdatePosition(t, 0, 0, west, 1, -1, 0)
}

func testAbs(t *testing.T, input, expected int) {
	actual := abs(input)

	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}

func TestAbs(t *testing.T) {
	testAbs(t, 0, 0)
	testAbs(t, 1, 1)
	testAbs(t, -1, 1)
}

func testFormatPosition(t *testing.T, x, y int, expected string) {
	actual := formatPosition(x, y)

	if actual != expected {
		t.Errorf("expected %s for (%d,%d), but got %s",
			expected, x, y, actual)
	}
}

func TestFormatPosition(t *testing.T) {
	testFormatPosition(t, 0, 0, "(0,0)")
	testFormatPosition(t, 1, 0, "(1,0)")
	testFormatPosition(t, 0, 1, "(0,1)")
	testFormatPosition(t, 1, 1, "(1,1)")
	testFormatPosition(t, 0, -1, "(0,-1)")
	testFormatPosition(t, -1, 0, "(-1,0)")
	testFormatPosition(t, -1, -1, "(-1,-1)")
}
