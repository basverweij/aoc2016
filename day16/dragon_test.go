package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDragonString(t *testing.T) {
	d := newDragon(1, 1)
	assert.Equal(t, "1", d.String())

	d = newDragon(0, 1)
	assert.Equal(t, "0", d.String())

	d = newDragon(4, 3)
	assert.Equal(t, "100", d.String())

	d = newDragon(1, 3)
	assert.Equal(t, "001", d.String())
}

func TestDragonExpand(t *testing.T) {
	d := newDragon(1, 1)
	d.expand()
	assert.Equal(t, &dragon{big.NewInt(4), 3}, d)

	d = newDragon(0, 1)
	d.expand()
	assert.Equal(t, &dragon{big.NewInt(1), 3}, d)

	d = newDragon(31, 5)
	d.expand()
	assert.Equal(t, &dragon{big.NewInt(1984), 11}, d)

	d = newDragon(3850, 12)
	d.expand()
	assert.Equal(t, &dragon{big.NewInt(31542000), 25}, d)
}
