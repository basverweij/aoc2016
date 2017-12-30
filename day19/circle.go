package main

type circle struct {
	head *position
	size int
}

type position struct {
	index       int
	numPresents int
	next        *position
}

func newCircle(size int) *circle {
	pos := &position{index: 1, numPresents: 1}
	c := &circle{head: pos, size: size}

	for i := 1; i < size; i++ {
		pos.next = &position{index: i + 1, numPresents: 1}
		pos = pos.next
	}

	pos.next = c.head

	return c
}

func (c *circle) removeAfter(pos *position) {
	if pos.next == c.head {
		c.head = pos.next.next
	}

	pos.next = pos.next.next
}

func (c *circle) take(from *position) {
	from.numPresents += from.next.numPresents

	c.removeAfter(from)

	c.size--
}

func (c *circle) done() bool {
	return c.size == 1
}

func (c *circle) takeAcross(from *position) {
	// move further half of the circle
	pos := from
	for i := c.size / 2; i > 1; i-- {
		pos = pos.next
	}

	from.numPresents += pos.next.numPresents

	c.removeAfter(pos)

	c.size--
}
