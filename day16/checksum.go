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

func partialChecksum(b *big.Int, length int) *big.Int {
	m := new(big.Int)
	m.Rsh(b, 1)
	b.Xor(b, m)

	m.SetInt64(1)
	m.Lsh(m, uint(length))
	m.Sub(m, big.NewInt(1))
	b.Xor(b, m)

	bw := b.Bits()
	if len(bw)%2 == 1 {
		bw = append(bw, 0)
	}

	cw := make([]big.Word, len(bw)/2)
	for i := 0; i < len(cw); i++ {
		cw[i] = compressWords(bw[i*2], bw[i*2+1])
	}

	c := new(big.Int)
	c.SetBits(cw)

	return c
}

func compressWords(w2, w1 big.Word) big.Word {
	return big.Word(
		compressedBytes[(w1>>56)&0xff]<<60 |
			compressedBytes[(w1>>48)&0xff]<<56 |
			compressedBytes[(w1>>40)&0xff]<<52 |
			compressedBytes[(w1>>32)&0xff]<<48 |
			compressedBytes[(w1>>24)&0xff]<<44 |
			compressedBytes[(w1>>16)&0xff]<<40 |
			compressedBytes[(w1>>8)&0xff]<<36 |
			compressedBytes[(w1)&0xff]<<32 |
			compressedBytes[(w2>>56)&0xff]<<28 |
			compressedBytes[(w2>>48)&0xff]<<24 |
			compressedBytes[(w2>>40)&0xff]<<20 |
			compressedBytes[(w2>>32)&0xff]<<16 |
			compressedBytes[(w2>>24)&0xff]<<12 |
			compressedBytes[(w2>>16)&0xff]<<8 |
			compressedBytes[(w2>>8)&0xff]<<4 |
			compressedBytes[(w2)&0xff])
}

var compressedBytes = [256]big.Word{
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	0,
	1,
	0,
	1,
	2,
	3,
	2,
	3,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	4,
	5,
	4,
	5,
	6,
	7,
	6,
	7,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	8,
	9,
	8,
	9,
	10,
	11,
	10,
	11,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
	12,
	13,
	12,
	13,
	14,
	15,
	14,
	15,
}
