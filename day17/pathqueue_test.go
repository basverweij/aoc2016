package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathQueuePushPop(t *testing.T) {
	q := new(pathQueue)
	assert.Nil(t, q.peek())
	assert.Nil(t, q.pop())

	q.push(&path{x: 1})
	q.push(&path{x: 2})
	q.push(&path{x: 3})

	assert.EqualValues(t, &path{x: 1}, q.peek())
	assert.EqualValues(t, &path{x: 1}, q.pop())
	assert.EqualValues(t, &path{x: 2}, q.peek())
	assert.EqualValues(t, &path{x: 2}, q.pop())
	assert.EqualValues(t, &path{x: 3}, q.peek())
	assert.EqualValues(t, &path{x: 3}, q.pop())
	assert.Nil(t, q.peek())
	assert.Nil(t, q.pop())
}
