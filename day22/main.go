package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	return len(viablePairs(input()))
}
