package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	findLow := 17
	findHigh := 61

	checkOutput1 := 0
	checkOutput2 := 1
	checkOutput3 := 2

	var inits []*initInstruction
	var cmps []*compareInstruction

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	b := bufio.NewReaderSize(f, 64)
	for {
		line, isPrefix, err := b.ReadLine()
		if line == nil || err == io.EOF {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or unexpected error")
		}

		init, cmp, err := parse(string(line))
		if err != nil {
			panic(err)
		}

		if init != nil {
			inits = append(inits, init)
		}

		if cmp != nil {
			cmps = append(cmps, cmp)
		}
	}

	m := newModel(inits, cmps)

	for {
		bots := m.step()
		if bots == nil {
			break
		}

		if lm, found := m.ValueCompares[findLow]; found {
			if botID, found := lm[findHigh]; found {
				fmt.Printf("bot #%d compares %d and %d\n", botID, findLow, findHigh)
				// break
			}
		}

		if o1, f1 := m.Outputs[checkOutput1]; f1 {
			if o2, f2 := m.Outputs[checkOutput2]; f2 {
				if o3, f3 := m.Outputs[checkOutput3]; f3 {
					fmt.Printf("outputs #%d, #%d, #%d contains %d, %d, %d; product = %d\n",
						checkOutput1, checkOutput2, checkOutput3,
						o1, o2, o3,
						o1*o2*o3)
					break
				}
			}
		}
	}
}
