package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	pwd := newPassword("abcde")
	assert.Len(t, pwd.value, 5)
	assert.Equal(t, "abcde", pwd.String())
}

func TestPasswordRotate(t *testing.T) {
	pwd := newPassword("abcde")

	pwd.rotate(true, 2)
	assert.Equal(t, "cdeab", pwd.String())

	pwd.rotate(false, 4)
	assert.Equal(t, "deabc", pwd.String())
}

func TestPasswordSwap(t *testing.T) {
	pwd := newPassword("abcde")

	pwd.swap(0, 4)
	assert.Equal(t, "ebcda", pwd.String())
}

func TestPasswordReverse(t *testing.T) {
	pwd := newPassword("abcde")

	pwd.swap(1, 3)
	assert.Equal(t, "adcbe", pwd.String())
}

func TestPasswordMove(t *testing.T) {
	pwd := newPassword("bcdea")

	pwd.move(1, 4)
	assert.Equal(t, "bdeac", pwd.String())

	pwd.move(3, 0)
	assert.Equal(t, "abdec", pwd.String())

	pwd.move(0, 4)
	assert.Equal(t, "bdeca", pwd.String())
}
