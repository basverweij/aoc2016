package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPositionPerform(t *testing.T) {
	pwd := newPassword("abcde")
	sp := &swapPosition{from: 4, to: 0}

	sp.perform(pwd)
	assert.Equal(t, "ebcda", pwd.String())

	sp.unperform(pwd)
	assert.Equal(t, "abcde", pwd.String())
}

func TestSwapPositionString(t *testing.T) {
	sp := &swapPosition{from: 4, to: 0}
	assert.Equal(t, "swap position 4 with position 0", sp.String())
}

func TestSwapLetterPerform(t *testing.T) {
	pwd := newPassword("ebcda")
	sl := &swapLetter{fromLetter: 'd', toLetter: 'e'}

	sl.perform(pwd)
	assert.Equal(t, "dbcea", pwd.String())

	sl.unperform(pwd)
	assert.Equal(t, "ebcda", pwd.String())
}

func TestSwapLetterString(t *testing.T) {
	sl := &swapLetter{fromLetter: 'd', toLetter: 'e'}
	assert.Equal(t, "swap letter d with letter e", sl.String())
}

func TestRotatePerform(t *testing.T) {
	r := &rotate{rotateLeft: true, steps: 1}
	pwd := newPassword("abcde")

	r.perform(pwd)
	assert.Equal(t, "bcdea", pwd.String())

	r.unperform(pwd)
	assert.Equal(t, "abcde", pwd.String())
}

func TestRotateString(t *testing.T) {
	r := &rotate{rotateLeft: true, steps: 1}
	assert.Equal(t, "rotate left 1 step", r.String())

	r = &rotate{rotateLeft: false, steps: 2}
	assert.Equal(t, "rotate right 2 steps", r.String())
}

func TestRotateLetterPerform(t *testing.T) {
	rl := &rotateLetter{letter: 'b'}
	pwd := newPassword("abdec")

	rl.perform(pwd)
	assert.Equal(t, "ecabd", pwd.String())

	rl.unperform(pwd)
	assert.Equal(t, "abdec", pwd.String())

	rl = &rotateLetter{letter: 'd'}
	pwd = newPassword("ecabd")

	rl.perform(pwd)
	assert.Equal(t, "decab", pwd.String())

	rl.unperform(pwd)
	assert.Equal(t, "ecabd", pwd.String())
}

func TestRotateLetterString(t *testing.T) {
	rl := &rotateLetter{letter: 'b'}
	assert.Equal(t, "rotate based on position of letter b", rl.String())
}

func TestReversePerform(t *testing.T) {
	rev := &reverse{from: 0, to: 4}
	pwd := newPassword("edcba")

	rev.perform(pwd)
	assert.Equal(t, "abcde", pwd.String())

	rev.unperform(pwd)
	assert.Equal(t, "edcba", pwd.String())
}

func TestReverseString(t *testing.T) {
	rev := &reverse{from: 0, to: 4}
	assert.Equal(t, "reverse positions 0 through 4", rev.String())
}

func TestMovePerform(t *testing.T) {
	mov := &move{from: 1, to: 4}
	pwd := newPassword("bcdea")

	mov.perform(pwd)
	assert.Equal(t, "bdeac", pwd.String())

	mov = &move{from: 3, to: 0}
	mov.perform(pwd)
	assert.Equal(t, "abdec", pwd.String())

	mov.unperform(pwd)
	assert.Equal(t, "bdeac", pwd.String())
}

func TestMoveString(t *testing.T) {
	mov := &move{from: 1, to: 4}
	assert.Equal(t, "move position 1 to position 4", mov.String())
}
