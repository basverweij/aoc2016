package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodePos(t *testing.T) {
	n := &node{"/dev/grid/node-x31-y29", 92, 66, 26, 71}
	x, y := n.pos()

	assert.Equal(t, 31, x)
	assert.Equal(t, 29, y)
}

func TestNodeMapAdjacentNodes(t *testing.T) {
	nodes := input()
	n := &node{"/dev/grid/node-x10-y10", 85, 72, 13, 84}

	assert.Equal(t, []*node{
		{"/dev/grid/node-x9-y10", 93, 69, 24, 74},
		{"/dev/grid/node-x11-y10", 93, 70, 23, 75},
		{"/dev/grid/node-x10-y9", 92, 72, 20, 78},
		{"/dev/grid/node-x10-y11", 93, 67, 26, 72},
	}, nodes.adjacentNodes(n))
}
