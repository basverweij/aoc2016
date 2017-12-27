package main

import (
	"fmt"
	"math/big"
)

func checksum(d *dragon, len int) string {
	c := new(big.Int)
	c.Set(d.val)

	// remove the bits :) we don't need
	c.Rsh(c, uint(d.len-len))

	for ; len%2 == 0; len /= 2 {
		c = partialChecksum(c, len)
	}

	return fmt.Sprintf(fmt.Sprintf("%%0%ds", len), c.Text(2))
}

func partialChecksum(b *big.Int, len int) *big.Int {
	m := new(big.Int)
	m.Rsh(b, 1)
	b.Xor(b, m)

	m.SetInt64(1)
	m.Lsh(m, uint(len))
	m.Sub(m, big.NewInt(1))
	b.Xor(b, m)

	c := new(big.Int)
	for i := 0; i < len; i += 2 {
		if i > 0 {
			c.Lsh(c, 1)
		}

		m.SetInt64(1)
		m.And(m, b)

		c.Or(c, m)

		b.Rsh(b, 2)
	}

	return reverse(c, len/2)
}
