package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())
}

func puzzle1() string {
	b, _ := ioutil.ReadFile("input.txt")

	// gfadhecb is not correct
	return scrambler("abcdefgh", input(string(b)))
}
