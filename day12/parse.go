package main

import (
	"strconv"
	"strings"
)

func parse(line string) instruction {
	parts := strings.Split(line, " ")

	switch parts[0] {
	case "cpy":
		if v, err := strconv.Atoi(parts[1]); err == nil {
			// copy value
			return &cpyVal{Val: v, To: parts[2]}
		}

		// copy register
		return &cpyReg{From: parts[1], To: parts[2]}
	case "jnz":
		if j, err := strconv.Atoi(parts[2]); err == nil {
			if v, err := strconv.Atoi(parts[1]); err == nil {
				// compare value
				return &jnzVal{Val: v, Jump: j}
			}

			// compare register
			return &jnzReg{Reg: parts[1], Jump: j}
		}

		panic("invalid jumps: " + parts[2])
	case "inc":
		return &inc{Reg: parts[1]}
	case "dec":
		return &dec{Reg: parts[1]}
	default:
		panic("invalid instruction: " + parts[0])
	}
}
