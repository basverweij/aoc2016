package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testHash = `abc`

func TestGenHash(t *testing.T) {
	for i := 0; i < 18; i++ {
		h := genHash(testHash, i)
		assert.Nil(t, h)
	}

	h := genHash(testHash, 18)
	assert.NotNil(t, h)
	assert.Equal(t, byte('8'), h.firstTriplet)

	h = genHash(testHash, 792)
	assert.NotNil(t, h)
	assert.Equal(t, byte('3'), h.firstTriplet)
	assert.Contains(t, h.quintets, byte('3'))

	h = genHash(testHash, 816)
	assert.NotNil(t, h)
	assert.Contains(t, h.quintets, byte('e'))
}

func TestGetSeries(t *testing.T) {
	assert.Equal(t, []byte{'a'}, getSeries("aaabcd", 3))
	assert.Equal(t, []byte{'a'}, getSeries("bcdaaa", 3))
	assert.Equal(t, []byte{'a'}, getSeries("bcdaaaxyz", 3))
	assert.Equal(t, []byte{'a', 'b'}, getSeries("bcdaaaxyzbbbpqr", 3))
	assert.Equal(t, []byte{'a'}, getSeries("aaaa", 3))

	assert.Nil(t, getSeries("aa", 3))
}
