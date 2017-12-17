package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const office10_7_10 = `  0123456789
0 .#.####.##
1 ..#..#...#
2 #....##...
3 ###.#.###.
4 .##..#..#.
5 ..##....#.
6 #...##.###
`

func TestOffice(t *testing.T) {
	o := newOffice(10, 7, 10)

	assert.Equal(t, office10_7_10, o.String())
}
