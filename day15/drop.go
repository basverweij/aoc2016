package main

func drop(discs []disc, tStart int) bool {
	for i, d := range discs {
		p := (d.startPos + tStart + i + 1) % d.positions

		if p != 0 {
			return false
		}
	}

	return true
}
