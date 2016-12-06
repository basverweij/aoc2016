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

	var msgs []string

	r := bufio.NewReaderSize(f, 12)
	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF || line == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or error during read")
		}

		msgs = append(msgs, string(line))
	}

	msg := decodeMsg(msgs)

	fmt.Printf("found message '%s'\n", msg)
}
