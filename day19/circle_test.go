package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircle(t *testing.T) {
	c := newCircle(3)
	assert.NotNil(t, c)

	assert.Equal(t, c.head, c.head.next.next.next)

	assert.Equal(t, 1, c.head.index)
	assert.Equal(t, 2, c.head.next.index)
	assert.Equal(t, 3, c.head.next.next.index)
}

func TestCircleTakeAndDone(t *testing.T) {
	c := newCircle(3)

	c.take(c.head.next)
	assert.Equal(t, c.head, c.head.next.next)

	c.take(c.head.next)
	assert.Equal(t, c.head, c.head.next)

	assert.True(t, c.done())

	assert.Equal(t, 2, c.head.index)
}
