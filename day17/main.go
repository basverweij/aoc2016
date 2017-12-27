package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() string {
	return shortestPath(input)
}

func puzzle2() int {
	return longestPath(input)
}
