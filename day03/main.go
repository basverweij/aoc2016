package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var inputPattern = regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var total, valid int

	// create buffer to hold sizes for each group of 3 rows
	sizesBuffer := make([][]int, 3)
	sizesBuffer[0] = make([]int, 3)
	sizesBuffer[1] = make([]int, 3)
	sizesBuffer[2] = make([]int, 3)

	b := bufio.NewReaderSize(f, 20)
	for {
		l, isPrefix, err := b.ReadLine()
		if err == io.EOF || l == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or error during read")
		}

		sizes := parseLine(string(l))
		for i := 0; i < 3; i++ {
			sizesBuffer[i][total%3] = sizes[i]
		}

		total++

		if total%3 == 0 {
			// process sizes buffer
			for i := 0; i < 3; i++ {
				if checkSizes(sizesBuffer[i]) {
					valid++
				}
			}
		}
	}

	if total%3 != 0 {
		panic(fmt.Sprintf("invalid number of rows: %d", total))
	}

	fmt.Printf("processed %d total lines, %d are valid\n", total, valid)
}

func parseLine(line string) []int {
	s := inputPattern.FindStringSubmatch(line)
	if s == nil || len(s) != 4 {
		panic("error parsing line: " + line)
	}

	sizes := make([]int, 3)
	for i := 0; i < 3; i++ {
		sizes[i], _ = strconv.Atoi(s[i+1])
	}

	return sizes
}

func checkSizes(sizes []int) bool {
	if sizes[0]+sizes[1] <= sizes[2] {
		return false
	}

	if sizes[0]+sizes[2] <= sizes[1] {
		return false
	}

	if sizes[1]+sizes[2] <= sizes[0] {
		return false
	}

	return true
}
