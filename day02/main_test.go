package main

import "testing"

func testUpdatePosition(t *testing.T, x, y int, dir direction, expectedX, expectedY int) {
	actualX, actualY := updatePosition(x, y, dir)

	if actualX != expectedX || actualY != expectedY {
		t.Errorf("expected (%d,%d) after updating (%d,%d) with %v, but got (%d,%d)",
			expectedX, expectedY, x, y, dir, actualX, actualY)
	}
}

func TestUpdatePosition(t *testing.T) {
	// "1"
	testUpdatePosition(t, 2, 0, up, 2, 0)
	testUpdatePosition(t, 2, 0, right, 2, 0)
	testUpdatePosition(t, 2, 0, down, 2, 1)
	testUpdatePosition(t, 2, 0, left, 2, 0)

	// "2"
	testUpdatePosition(t, 1, 1, up, 1, 1)
	testUpdatePosition(t, 1, 1, right, 2, 1)
	testUpdatePosition(t, 1, 1, down, 1, 2)
	testUpdatePosition(t, 1, 1, left, 1, 1)

	// "4"
	testUpdatePosition(t, 3, 1, up, 3, 1)
	testUpdatePosition(t, 3, 1, right, 3, 1)
	testUpdatePosition(t, 3, 1, down, 3, 2)
	testUpdatePosition(t, 3, 1, left, 2, 1)

	// "5"
	testUpdatePosition(t, 0, 2, up, 0, 2)
	testUpdatePosition(t, 0, 2, right, 1, 2)
	testUpdatePosition(t, 0, 2, down, 0, 2)
	testUpdatePosition(t, 0, 2, left, 0, 2)

	// "9"
	testUpdatePosition(t, 4, 2, up, 4, 2)
	testUpdatePosition(t, 4, 2, right, 4, 2)
	testUpdatePosition(t, 4, 2, down, 4, 2)
	testUpdatePosition(t, 4, 2, left, 3, 2)

	// "A"
	testUpdatePosition(t, 1, 3, up, 1, 2)
	testUpdatePosition(t, 1, 3, right, 2, 3)
	testUpdatePosition(t, 1, 3, down, 1, 3)
	testUpdatePosition(t, 1, 3, left, 1, 3)

	// "C"
	testUpdatePosition(t, 3, 3, up, 3, 2)
	testUpdatePosition(t, 3, 3, right, 3, 3)
	testUpdatePosition(t, 3, 3, down, 3, 3)
	testUpdatePosition(t, 3, 3, left, 2, 3)

	// "D"
	testUpdatePosition(t, 2, 4, up, 2, 3)
	testUpdatePosition(t, 2, 4, right, 2, 4)
	testUpdatePosition(t, 2, 4, down, 2, 4)
	testUpdatePosition(t, 2, 4, left, 2, 4)
}

func testPositionToDigit(t *testing.T, x, y int, expected string) {
	actual := positionToDigit(x, y)

	if actual != expected {
		t.Errorf("epxected %s for (%d,%d), but got %s",
			expected, x, y, actual)
	}
}

func TestPositionToDigit(t *testing.T) {
	testPositionToDigit(t, 2, 0, "1")
	testPositionToDigit(t, 1, 1, "2")
	testPositionToDigit(t, 2, 1, "3")
	testPositionToDigit(t, 3, 1, "4")
	testPositionToDigit(t, 0, 2, "5")
	testPositionToDigit(t, 1, 2, "6")
	testPositionToDigit(t, 2, 2, "7")
	testPositionToDigit(t, 3, 2, "8")
	testPositionToDigit(t, 4, 2, "9")
	testPositionToDigit(t, 1, 3, "A")
	testPositionToDigit(t, 2, 3, "B")
	testPositionToDigit(t, 3, 3, "C")
	testPositionToDigit(t, 2, 4, "D")
}
