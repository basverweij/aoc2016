package main

import "testing"

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

	assertEquals(t, office10_7_10, o.String(), "office 10x7#10")
}
