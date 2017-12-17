package main

var nibbleBits = []uint8{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

func numBits(x uint64) uint8 {
	var n uint8
	for ; x != 0; x >>= 4 {
		n += nibbleBits[x&0xf]
	}

	return n
}
