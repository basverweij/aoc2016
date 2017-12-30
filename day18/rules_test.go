package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyRules(t *testing.T) {
	row := []tile{safe, safe, trap, trap, safe}

	next := applyRules(row)
	assert.Equal(t, []tile{safe, trap, trap, trap, trap}, next)

	next = applyRules(next)
	assert.Equal(t, []tile{trap, trap, safe, safe, trap}, next)
}
