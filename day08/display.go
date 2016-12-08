package main

import (
	"strings"
)

type display struct {
	Pixels [][]bool
}

func newDisplay(width, height int) *display {
	d := &display{Pixels: make([][]bool, height)}

	for y := 0; y < height; y++ {
		d.Pixels[y] = make([]bool, width)
	}

	return d
}

func (d *display) String() string {
	h := len(d.Pixels)
	w := len(d.Pixels[0])
	rows := make([]string, h)

	for y := 0; y < h; y++ {
		buf := make([]byte, w)
		for x := 0; x < w; x++ {
			if d.Pixels[y][x] {
				buf[x] = '#'
			} else {
				buf[x] = '.'
			}
		}

		rows[y] = string(buf)
	}

	return strings.Join(rows, "\n")
}

func (d *display) ActivePixelCount() int {
	h := len(d.Pixels)
	w := len(d.Pixels[0])
	n := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if d.Pixels[y][x] {
				n++
			}
		}
	}

	return n
}

type instruction interface {
	Execute(*display)
}
