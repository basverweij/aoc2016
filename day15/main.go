package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	for tStart := 0; ; tStart++ {
		if drop(input, tStart) {
			return tStart
		}
	}
}
