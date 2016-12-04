package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

var directionMapping map[rune]direction

func init() {
	directionMapping = make(map[rune]direction)
	directionMapping['U'] = up
	directionMapping['R'] = right
	directionMapping['D'] = down
	directionMapping['L'] = left
}

var keypad = [][]rune{
	{' ', ' ', '1', ' ', ' '},
	{' ', '2', '3', '4', ' '},
	{'5', '6', '7', '8', '9'},
	{' ', 'A', 'B', 'C', ' '},
	{' ', ' ', 'D', ' ', ' '},
}

func main() {
	var x, y int = 0, 2

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	code := ""

	r := bufio.NewReaderSize(f, 8192)
	for {
		l, isPrefix, err := r.ReadLine()
		if err == io.EOF || l == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or error during read")
		}

		x, y = processLine(x, y, string(l))

		code += positionToDigit(x, y)
	}

	fmt.Printf("the code is: %s\n", code)
}

func processLine(x, y int, line string) (int, int) {
	for _, r := range line {
		dir, valid := directionMapping[r]
		if !valid {
			panic("invalid direction: " + string(r))
		}

		x, y = updatePosition(x, y, dir)
	}

	return x, y
}

func updatePosition(x, y int, dir direction) (int, int) {
	if dir == up && y > 0 && keypad[y-1][x] != ' ' {
		return x, y - 1
	} else if dir == right && x < len(keypad[y])-1 && keypad[y][x+1] != ' ' {
		return x + 1, y
	} else if dir == down && y < len(keypad)-1 && keypad[y+1][x] != ' ' {
		return x, y + 1
	} else if dir == left && x > 0 && keypad[y][x-1] != ' ' {
		return x - 1, y
	}

	return x, y
}

func positionToDigit(x, y int) string {
	return string(keypad[y][x])
}
