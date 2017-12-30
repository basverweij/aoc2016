package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	return totalSafeTiles(input, 40)
}

func puzzle2() int {
	return totalSafeTiles(input, 400000)
}
