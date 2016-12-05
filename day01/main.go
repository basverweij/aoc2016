package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var instructionPattern = regexp.MustCompile("([RL])([0-9]+)")

type heading int

const (
	north heading = iota
	east
	south
	west
)

type direction int

const (
	right direction = iota
	left
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	s := string(b)
	instructions := strings.Split(s, ", ")

	fmt.Printf("found %d instruction(s)\n", len(instructions))

	var x, y int
	var currentHeading heading

	// keep set of all old positions
	oldPositions := make(map[string]bool)
	oldPositions[formatPosition(x, y)] = true

	foundHQ := false

	for _, instruction := range instructions {
		dir, num, err := parseInstruction(instruction)
		if err != nil {
			panic(err)
		}

		currentHeading = updateHeading(currentHeading, dir)

		for i := 0; i < num; i++ {
			x, y = updatePosition(x, y, currentHeading, 1)

			f := formatPosition(x, y)
			if _, exists := oldPositions[f]; exists {
				// already been here once
				foundHQ = true
				break
			}

			oldPositions[f] = true
		}

		if foundHQ {
			break
		}
	}

	fmt.Printf("arrived at (%d,%d), taxicab distance %d\n", x, y, abs(x)+abs(y))
}

func parseInstruction(instruction string) (dir direction, num int, err error) {
	parts := instructionPattern.FindStringSubmatch(instruction)
	if parts == nil || len(parts) != 3 {
		err = fmt.Errorf("parsing instruction: %s", instruction)
		return
	}

	if parts[1] == "R" {
		dir = right
	} else {
		dir = left
	}

	num, err = strconv.Atoi(parts[2])
	if err != nil {
		return
	}

	return
}

func updateHeading(currentHeading heading, dir direction) heading {
	i := int(currentHeading)

	if dir == right {
		i++
	} else {
		i--
	}

	i = (i + 4) % 4

	return heading(i)
}

func updatePosition(x, y int, currentHeading heading, num int) (int, int) {
	switch currentHeading {
	case north:
		return x, y + num
	case east:
		return x + num, y
	case south:
		return x, y - num
	case west:
		return x - num, y
	}

	panic("invalid heading")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func formatPosition(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}
