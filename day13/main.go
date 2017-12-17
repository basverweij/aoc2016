package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	o := newOffice(80, 80, input)
	return shortestPath(o, pos{1, 1}, pos{31, 39})
}
