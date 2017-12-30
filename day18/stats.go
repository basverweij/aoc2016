package main

func safeTiles(row []tile) int {
	n := 0

	for _, tile := range row {
		if tile == safe {
			n++
		}
	}

	return n
}

func totalSafeTiles(row []tile, rows int) int {
	n := safeTiles(row)

	for i := 1; i < rows; i++ {
		row = applyRules(row)
		n += safeTiles(row)
	}

	return n
}
