package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	keys := genKeys(inputSalt, 64, 0)
	return keys[63].index
}

func puzzle2() int {
	keys := genKeys(inputSalt, 64, inputStretching)
	return keys[63].index
}
