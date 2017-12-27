package main

import (
	"fmt"
	"math/big"
)

type dragon struct {
	val *big.Int
	len int
}

func newDragon(start int64, len int) *dragon {
	return &dragon{big.NewInt(start), len}
}

func (d *dragon) String() string {
	return fmt.Sprintf(fmt.Sprintf("%%0%ds", d.len), d.val.Text(2))
}

func (d *dragon) expand() {
	b := reverse(d.val, d.len)

	m := big.NewInt(1)
	m.Lsh(m, uint(d.len))
	m.Sub(m, big.NewInt(1))
	b.Xor(b, m)

	d.val.Lsh(d.val, uint(d.len+1))
	d.val.Or(d.val, b)

	d.len = 2*d.len + 1
}
