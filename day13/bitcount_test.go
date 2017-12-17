package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitCount(t *testing.T) {
	assert.Equal(t, uint8(0), numBits(0))
	assert.Equal(t, uint8(1), numBits(1))
	assert.Equal(t, uint8(32), numBits(0x5a5a5a5a5a5a5a5a))
	assert.Equal(t, uint8(64), numBits(0xffffffffffffffff))
}
