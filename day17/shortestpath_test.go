package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	assert.Equal(t, "", shortestPath("hijkl"))
	assert.Equal(t, "DDRRRD", shortestPath("ihgpwlah"))
	assert.Equal(t, "DDUDRLRRUDRD", shortestPath("kglvqrro"))
	assert.Equal(t, "DRURDRUDDLLDLUURRDULRLDUUDDDRR", shortestPath("ulqzkmiv"))
}
