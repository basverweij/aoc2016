package main

import (
	"strconv"
)

func isWall(x, y, design int) bool {
	var z = x*x + 3*x + 2*x*y + y + y*y
	var n = numBits(uint64(z + design))

	return n%2 == 1
}

type office [][]bool

func newOffice(sizeX, sizeY, design int) office {
	o := make([][]bool, sizeY)

	for y := 0; y < sizeY; y++ {
		o[y] = make([]bool, sizeX)

		for x := 0; x < sizeX; x++ {
			o[y][x] = isWall(x, y, design)
		}
	}

	return o
}

func (o office) String() string {
	sizeY := len(o)
	sizeX := len(o[0])

	s := "  "
	for x := 0; x < sizeX; x++ {
		s += strconv.Itoa(x % 10)
	}
	s += "\n"

	for y := 0; y < sizeY; y++ {
		s += strconv.Itoa(y%10) + " "

		for x := 0; x < sizeX; x++ {
			if o[y][x] {
				s += "#"
			} else {
				s += "."
			}
		}

		s += "\n"
	}

	return s
}
