package main

func playWhiteElephant(numElves int) int {
	c := newCircle(numElves)

	for pos := c.head; !c.done(); pos = pos.next {
		c.take(pos)
	}

	return c.head.index
}

func playWhiteElephantAcross(numElves int) int {
	if numElves < 4 {
		panic("invalid number of elves")
	}

	singles := true
	size := 3
	offset := 3

	for offset+size < numElves {
		offset += size

		if singles {
			singles = false
		} else {
			singles = true
			size *= 3
		}
	}

	if singles {
		return numElves - offset
	}

	return (numElves-offset)*2 + size
}
