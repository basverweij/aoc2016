package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenKeys(t *testing.T) {
	keys := genKeys("abc", 64)
	assert.Len(t, keys, 64)
}
