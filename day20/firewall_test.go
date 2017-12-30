package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLowestAllowed(t *testing.T) {
	assert.Equal(t, uint32(3), findLowestAllowed(testInput))
}

func TestFindAmountAllowed(t *testing.T) {
	assert.Equal(t, uint32(2), findAmountAllowed(testInput, 9))
}
