package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLowestAllowed(t *testing.T) {
	assert.Equal(t, uint32(3), findLowestAllowed(testInput))
}
