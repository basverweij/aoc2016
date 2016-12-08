package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var rectPattern = regexp.MustCompile(`rect (\d+)x(\d+)`)
var rotateRowPattern = regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
var rotateColumnPattern = regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)

func parse(s string) (instruction, error) {
	if m := rectPattern.FindStringSubmatch(s); m != nil {
		w, _ := strconv.Atoi(m[1])
		h, _ := strconv.Atoi(m[2])

		return &rectInstruction{Width: w, Height: h}, nil
	}

	if m := rotateRowPattern.FindStringSubmatch(s); m != nil {
		y, _ := strconv.Atoi(m[1])
		n, _ := strconv.Atoi(m[2])

		return &rotateRowInstruction{Y: y, N: n}, nil
	}

	if m := rotateColumnPattern.FindStringSubmatch(s); m != nil {
		x, _ := strconv.Atoi(m[1])
		n, _ := strconv.Atoi(m[2])

		return &rotateColumnInstruction{X: x, N: n}, nil
	}

	return nil, fmt.Errorf("invalid input: %s", s)
}
