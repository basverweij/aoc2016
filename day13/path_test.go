package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestPath(t *testing.T) {
	o := newOffice(10, 7, 10)

	assert.Equal(t, 11, shortestPath(o, pos{1, 1}, findShortesPathTo(pos{7, 4})))
}

func TestFindUniqueLocationsAtLen(t *testing.T) {
	o := newOffice(10, 7, 10)

	assert.Equal(t, 6, shortestPath(o, pos{1, 1}, findUniqueLocationsAtLen(3)))
}
