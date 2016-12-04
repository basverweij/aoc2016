package main

func decodeRoomID(roomID string, sectorID int) string {
	rs := make([]rune, len(roomID))
	for i, r := range roomID {
		if r == '-' {
			rs[i] = ' '
			continue
		}

		rs[i] = shiftRune(r, sectorID)
	}

	return string(rs)
}

func shiftRune(r rune, sectorID int) rune {
	return rune('a' + ((int(r-'a') + sectorID) % 26))
}
