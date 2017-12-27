package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	for tStart := 0; tStart < 5; tStart++ {
		assert.False(t, drop(testInput, tStart))
	}

	assert.True(t, drop(testInput, 5))
}
