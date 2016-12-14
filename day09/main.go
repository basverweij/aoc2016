package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	s, err := decompress(string(b))
	if err != nil {
		panic(err)
	}

	fmt.Printf("decompressed length is: %d\n", len([]rune(s)))
}
