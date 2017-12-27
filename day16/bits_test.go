package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	b := big.NewInt(0x1200340056007800)

	r := reverse(b, 62)

	assert.Equal(t, big.NewInt(0x7801a800b0012), r)

	r = reverse(r, 62)

	assert.Equal(t, big.NewInt(0x1200340056007800), r)
}

func TestBitCount(t *testing.T) {
	b := new(big.Int)
	assert.Equal(t, 0, bitCount(b))

	b = big.NewInt(0x7fffffffffffffff)
	assert.Equal(t, 63, bitCount(b))
}
