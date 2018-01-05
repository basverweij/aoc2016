package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	assert.Len(t, input(), 960)
}
