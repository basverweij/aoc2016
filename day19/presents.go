package main

func playWhiteElephant(numElves int) int {
	c := newCircle(numElves)

	for pos := c.head; !c.done(); pos = pos.next {
		c.take(pos)
	}

	return c.head.index
}
