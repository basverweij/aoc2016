package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrambler(t *testing.T) {
	ops := input(testInput)

	// remove rotate right
	ops = ops[0:8]

	assert.Equal(t, "decab", scrambler("abcde", ops))
}
