package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	assert.Equal(t, "100", checksum(newDragon(6504, 13), 12))
}

func TestCompressWords(t *testing.T) {
	assert.Equal(t,
		big.Word(0x0000000100000001),
		compressWords(0x1, 0x1))

	assert.Equal(t,
		big.Word(0x1111111111111111),
		compressWords(0x0101010101010101, 0x0101010101010101))
}
