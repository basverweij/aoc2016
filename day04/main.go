package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var roomPattern = regexp.MustCompile(`([a-z-]+)-(\d+)\[([a-z]{1,5})\]`)

func main() {
	fmt.Println("running aoc2016-day04")

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	var validRooms []string
	var sectorIDs []int

	b := bufio.NewReaderSize(f, 128)
	for {
		line, isPrefix, err := b.ReadLine()
		if err == io.EOF || line == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or error during read")
		}

		isValid, sectorID, roomID, err := parseLine(string(line))
		if err != nil {
			panic(err)
		}

		if !isValid {
			continue
		}

		sum += sectorID

		validRooms = append(validRooms, decodeRoomID(roomID, sectorID))
		sectorIDs = append(sectorIDs, sectorID)
	}

	fmt.Printf("sector id sum = %d\n", sum)
	fmt.Printf("found %d valid rooms\n", len(validRooms))

	for i, room := range validRooms {
		if !strings.Contains(room, "north") {
			continue
		}

		fmt.Printf("#%03d: %s\n", sectorIDs[i], room)
	}
}

func parseLine(line string) (isValid bool, sectorID int, roomID string, err error) {
	m := roomPattern.FindStringSubmatch(line)
	if m == nil || len(m) != 4 {
		err = fmt.Errorf("unable to parse line: %s", line)
		return
	}

	roomID = m[1]

	cs := calcChecksum(roomID)
	if cs != m[3] {
		return
	}

	sectorID, err = strconv.Atoi(m[2])
	if err != nil {
		err = fmt.Errorf("invalid section id: %s", m[2])
	}

	isValid = true

	return
}
