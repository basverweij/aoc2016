package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	for tStart := 0; ; tStart++ {
		if drop(input, tStart) {
			return tStart
		}
	}
}

func puzzle2() int {
	for tStart := 0; ; tStart++ {
		if drop(input2, tStart) {
			return tStart
		}
	}
}
