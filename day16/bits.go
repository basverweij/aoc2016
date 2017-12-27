package main

import (
	"math/big"
	"math/bits"
)

func reverse(b *big.Int, size int) *big.Int {
	w := b.Bits()

	rw := make([]big.Word, len(w))
	for i := 0; i < len(w); i++ {
		rw[len(w)-i-1] = big.Word(bits.Reverse(uint(w[i])))
	}

	r := new(big.Int)
	r.SetBits(rw)

	sh := uint(64 - size%64)
	if sh > 0 {
		r.Rsh(r, sh)
	}

	return r
}

func bitCount(b *big.Int) int {
	count := 0
	for _, w := range b.Bits() {
		count += bits.OnesCount64(uint64(w))
	}

	return count
}
