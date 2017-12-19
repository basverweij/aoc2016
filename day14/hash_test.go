package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSalt = `abc`

func TestGenHash(t *testing.T) {
	for i := 0; i < 18; i++ {
		h := genHash(testSalt, i, 0)
		assert.Nil(t, h)
	}

	h := genHash(testSalt, 18, 0)
	assert.NotNil(t, h)
	assert.Equal(t, byte('8'), h.firstTriplet)

	h = genHash(testSalt, 792, 0)
	assert.NotNil(t, h)
	assert.Equal(t, byte('3'), h.firstTriplet)
	assert.Contains(t, h.quintets, byte('3'))

	h = genHash(testSalt, 816, 0)
	assert.NotNil(t, h)
	assert.Contains(t, h.quintets, byte('e'))
}

func TestGenHashWithStretching(t *testing.T) {
	h := genHash(testSalt, 5, 2016)
	assert.Equal(t, byte('2'), h.firstTriplet)

	h = genHash(testSalt, 10, 2016)
	assert.Equal(t, byte('e'), h.firstTriplet)

	h = genHash(testSalt, 89, 2016)
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
