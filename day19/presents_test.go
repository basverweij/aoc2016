package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayWhiteElephant(t *testing.T) {
	assert.Equal(t, 3, playWhiteElephant(5))
}

func TestPlayWhiteElephantAcross(t *testing.T) {
	assert.Equal(t, 2, playWhiteElephantAcross(5))
	assert.Equal(t, 81, playWhiteElephantAcross(81))
	assert.Equal(t, 1, playWhiteElephantAcross(82))
	assert.Equal(t, 8, playWhiteElephantAcross(89))
}
