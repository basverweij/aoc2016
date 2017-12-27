package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())

	fmt.Printf("Puzzle 2: %s\n", puzzle2())
}

func puzzle1() string {
	return gen(272, inputValue, inputLen)
}

func puzzle2() string {
	return gen(35651584, inputValue, inputLen)
}
