package main

import "testing"

func TestBitCount(t *testing.T) {
	assertEquals(t, uint8(0), numBits(0), "0")
	assertEquals(t, uint8(1), numBits(1), "1")
	assertEquals(t, uint8(32), numBits(0x5a5a5a5a5a5a5a5a), "0x5a5a5a5a5a5a5a5a")
	assertEquals(t, uint8(64), numBits(0xffffffffffffffff), "0xffffffffffffffff")
}
