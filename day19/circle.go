package main

type circle struct {
	head *position
}

type position struct {
	index       int
	numPresents int
	next        *position
}

func newCircle(size int) *circle {
	pos := &position{index: 1, numPresents: 1}
	c := &circle{head: pos}

	for i := 1; i < size; i++ {
		pos.next = &position{index: i + 1, numPresents: 1}
		pos = pos.next
	}

	pos.next = c.head

	return c
}

func (c *circle) take(from *position) {
	from.numPresents += from.next.numPresents

	if from.next == c.head {
		c.head = from.next.next
	}

	from.next = from.next.next
}

func (c *circle) done() bool {
	return c.head == c.head.next
}
