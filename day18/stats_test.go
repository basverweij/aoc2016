package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeTiles(t *testing.T) {
	row := []tile{}
	assert.Equal(t, 0, safeTiles(row))

	row = []tile{trap}
	assert.Equal(t, 0, safeTiles(row))

	row = []tile{safe}
	assert.Equal(t, 1, safeTiles(row))

	row = []tile{safe, safe, safe, safe}
	assert.Equal(t, 4, safeTiles(row))
}

func TestTotalSafeTiles(t *testing.T) {
	row := []tile{safe, trap, trap, safe, trap, safe, trap, trap, trap, trap}
	assert.Equal(t, 38, totalSafeTiles(row, 10))
}
