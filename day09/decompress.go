package main

import (
	"fmt"
	"strconv"
)

func decompress(s string) (string, error) {
	var buf []rune

	r := []rune(s)
	n := len(r)
	for i := 0; i < n; i++ {
		runeValue := r[i]

		// check for start of marker
		if runeValue == '(' {
			// scan int until 'x'
			markerLen, j, err := scanInt(r, i+1, 'x')
			if err != nil {
				return "", err
			}

			// scan int until ')'
			markerCount, j, err := scanInt(r, j+1, ')')
			if err != nil {
				return "", err
			}

			if j+1+markerLen > len(r) {
				return "", fmt.Errorf("invalid marker: %s", string(r[i:]))
			}

			for ; markerCount > 0; markerCount-- {
				buf = append(buf, r[j+1:j+1+markerLen]...)
			}

			// move i to last position of repeated section
			i = j + markerLen

			continue
		}

		buf = append(buf, runeValue)
	}

	return string(buf), nil
}

// scanInt searches for an int in the rune slice, starting at 'start'
// It returns the value of the int if found, the end position if found,
// or an error in case the end of the slice is reached, or when the sub slice
// does not hold a valid int.
func scanInt(r []rune, start int, stop rune) (value int, end int, err error) {
	for end = start; end < len(r); end++ {
		if r[end] == stop {
			break
		}
	}

	if end >= len(r) || r[end] != stop {
		err = fmt.Errorf("stop rune not found")
		return
	}

	value, err = strconv.Atoi(string(r[start:end]))
	return
}
