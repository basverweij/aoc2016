package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var roomPattern = regexp.MustCompile(`([a-z-]+)-(\d+)\[([a-z]{1,5})\]`)

func main() {
	fmt.Println("running aoc2016-day04")

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	b := bufio.NewReaderSize(f, 128)
	for {
		line, isPrefix, err := b.ReadLine()
		if err == io.EOF || line == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or error during read")
		}

		isValid, sectionID, err := parseLine(string(line))
		if err != nil {
			panic(err)
		}

		if !isValid {
			continue
		}

		sum += sectionID
	}

	fmt.Printf("section id sum = %d\n", sum)
}

func parseLine(line string) (isValid bool, sectionID int, err error) {
	m := roomPattern.FindStringSubmatch(line)
	if m == nil || len(m) != 4 {
		err = fmt.Errorf("unable to parse line: %s", line)
		return
	}

	cs := calcChecksum(m[1])
	if cs != m[3] {
		return
	}

	sectionID, err = strconv.Atoi(m[2])
	if err != nil {
		err = fmt.Errorf("invalid section id: %s", m[2])
	}

	isValid = true

	return
}
