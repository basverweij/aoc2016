package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	o := newOffice(80, 80, input)

	return shortestPath(o, pos{1, 1}, findShortesPathTo(pos{31, 39}))
}

func puzzle2() int {
	o := newOffice(80, 80, input)

	return shortestPath(o, pos{1, 1}, findUniqueLocationsAtLen(50))
}
