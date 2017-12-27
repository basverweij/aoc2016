package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	assert.Equal(t, "100", checksum(newDragon(6504, 13), 12))
}
