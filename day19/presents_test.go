package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayWhiteElephant(t *testing.T) {
	assert.Equal(t, 3, playWhiteElephant(5))
}
