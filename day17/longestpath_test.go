package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPath(t *testing.T) {
	assert.Equal(t, 0, longestPath("hijkl"))
	assert.Equal(t, 370, longestPath("ihgpwlah"))
	assert.Equal(t, 492, longestPath("kglvqrro"))
	assert.Equal(t, 830, longestPath("ulqzkmiv"))
}
