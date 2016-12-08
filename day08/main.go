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

	d := newDisplay(50, 6)

	b := bufio.NewReaderSize(f, 32)
	for {
		line, isPrefix, err := b.ReadLine()
		if err == io.EOF || line == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or unexpected error")
		}

		i, err := parse(string(line))
		if err != nil {
			panic(err)
		}

		i.Execute(d)
	}

	fmt.Println(d.String())

	fmt.Printf("display has %d active pixels\n", d.ActivePixelCount())
}
