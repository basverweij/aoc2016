package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())

	fmt.Printf("Puzzle 2: %s\n", puzzle2())
}

func puzzle1() string {
	b, _ := ioutil.ReadFile("input.txt")

	return scrambler("abcdefgh", input(string(b)))
}

func puzzle2() string {
	b, _ := ioutil.ReadFile("input.txt")

	return unscrambler("fbgdceah", input(string(b)))
}
