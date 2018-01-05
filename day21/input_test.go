package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
rotate right 2 steps`

func TestInput(t *testing.T) {
	ops := input(testInput)
	assert.Len(t, ops, 9)

	assert.Equal(t, []operation{
		&swapPosition{4, 0},
		&swapLetter{'d', 'b'},
		&reverse{0, 4},
		&rotate{true, 1},
		&move{1, 4},
		&move{3, 0},
		&rotateLetter{'b'},
		&rotateLetter{'d'},
		&rotate{false, 2},
	}, ops)
}
