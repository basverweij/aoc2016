package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	// create system with four registers
	s := newSystem("a", "b", "c", "d")
	s.Reg["c"] = 1

	// parse input file
	b := bufio.NewReaderSize(f, 16)
	for {
		line, isPrefix, err := b.ReadLine()
		if line == nil || err == io.EOF {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or unexpected error")
		}

		s.addInstruction(parse(string(line)))
	}

	fmt.Println(s.instructions)

	// run system
	s.Run()

	fmt.Printf("register 'a' contains %d\n", s.Reg["a"])
}
