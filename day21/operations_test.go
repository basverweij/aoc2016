package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformSwapPosition(t *testing.T) {
	pwd := newPassword("abcde")
	sp := &swapPosition{from: 4, to: 0}

	sp.perform(pwd)
	assert.Equal(t, "ebcda", pwd.String())
}

func TestPerformSwapLetter(t *testing.T) {
	pwd := newPassword("ebcda")
	sl := &swapLetter{fromLetter: 'd', toLetter: 'b'}

	sl.perform(pwd)
	assert.Equal(t, "edcba", pwd.String())
}

func TestPerformRotate(t *testing.T) {
	r := &rotate{rotateLeft: true, steps: 1}
	pwd := newPassword("abcde")

	r.perform(pwd)
	assert.Equal(t, "bcdea", pwd.String())
}

func TestPerformRotateLetter(t *testing.T) {
	rl := &rotateLetter{letter: 'b'}
	pwd := newPassword("abdec")

	rl.perform(pwd)
	assert.Equal(t, "ecabd", pwd.String())

	rl = &rotateLetter{letter: 'd'}

	rl.perform(pwd)
	assert.Equal(t, "decab", pwd.String())
}

func TestPerformReverse(t *testing.T) {
	rev := &reverse{from: 0, to: 4}
	pwd := newPassword("edcba")

	rev.perform(pwd)
	assert.Equal(t, "abcde", pwd.String())
}

func TestPerformMove(t *testing.T) {
	mov := &move{from: 1, to: 4}
	pwd := newPassword("bcdea")

	mov.perform(pwd)
	assert.Equal(t, "bdeac", pwd.String())

	mov = &move{from: 3, to: 0}
	mov.perform(pwd)
	assert.Equal(t, "abdec", pwd.String())
}
