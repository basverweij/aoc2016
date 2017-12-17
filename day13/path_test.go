package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	o := newOffice(10, 7, 10)

	assert.Equal(t, 11, shortestPath(o, pos{1, 1}, pos{7, 4}))
}
