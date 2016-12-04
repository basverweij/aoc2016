package main

import (
	"fmt"
	"testing"
)

func (f *runeFreq) String() string {
	return fmt.Sprintf("%s#%d", string(f.Rune), f.Freq)
}

func testCalcChecksum(t *testing.T, roomID, expectedChecksum string) {
	actualChecksum := calcChecksum(roomID)
	if actualChecksum != expectedChecksum {
		t.Errorf("invalid checksum for room %s: expected %s, but got %s",
			roomID, expectedChecksum, actualChecksum)
	}
}

func TestCalcChecksum(t *testing.T) {
	testCalcChecksum(t, "aaaaa-bbb-z-y-x", "abxyz")
	testCalcChecksum(t, "a-b-c-d-e-f-g-h", "abcde")
	testCalcChecksum(t, "not-a-real-room", "oarel")
}
