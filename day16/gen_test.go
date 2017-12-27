package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	c := gen(20, 16, 5)
	assert.Equal(t, "01100", c)
}
