package main

import "testing"

func testDecodeRoomID(t *testing.T, roomID string, sectorID int, expectedName string) {
	actualName := decodeRoomID(roomID, sectorID)
	if actualName != expectedName {
		t.Errorf("expected %s for room id %s and sector id %d, but got %s",
			expectedName, roomID, sectorID, actualName)
	}
}

func TestDecodeRoomID(t *testing.T) {
	testDecodeRoomID(t, "qzmt-zixmtkozy-ivhz", 343, "very encrypted name")
}
