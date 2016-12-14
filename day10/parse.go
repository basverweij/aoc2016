package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type initInstruction struct {
	Bot   int
	Value int
}

type compareInstruction struct {
	Bot       int
	Low       int
	LowIsBot  bool
	High      int
	HighIsBot bool
}

var (
	initInstructionPattern    = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	compareInstructionPattern = regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)
)

// parse process the input line and returns either an init instruction,
// a compare instruction, or an error
func parse(line string) (*initInstruction, *compareInstruction, error) {
	m := initInstructionPattern.FindStringSubmatch(line)
	if m != nil && len(m) == 3 {
		init := &initInstruction{}
		init.Value, _ = strconv.Atoi(m[1])
		init.Bot, _ = strconv.Atoi(m[2])

		return init, nil, nil
	}

	m = compareInstructionPattern.FindStringSubmatch(line)
	if m != nil && len(m) == 6 {
		cmp := &compareInstruction{}
		cmp.Bot, _ = strconv.Atoi(m[1])
		cmp.LowIsBot = m[2] == "bot"
		cmp.Low, _ = strconv.Atoi(m[3])
		cmp.HighIsBot = m[4] == "bot"
		cmp.High, _ = strconv.Atoi(m[5])

		return nil, cmp, nil
	}

	return nil, nil, fmt.Errorf("invalid line: %s", line)
}
