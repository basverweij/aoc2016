package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() uint32 {
	return findLowestAllowed(input)
}

func puzzle2() uint32 {
	// 4267919 is too high
	return findAmountAllowed(input, math.MaxUint32)
}
