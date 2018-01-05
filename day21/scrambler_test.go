package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrambler(t *testing.T) {
	ops := input(testInput)
	assert.Equal(t, "abdec", scrambler("abcde", ops))
}

func TestUnscrambler(t *testing.T) {
	ops := input(testInput)
	assert.Equal(t, "abcde", unscrambler("abdec", ops))
}
