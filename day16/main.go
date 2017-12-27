package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())
}

func puzzle1() string {
	return gen(272, inputValue, inputLen)
}
