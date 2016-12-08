package main

type rectInstruction struct {
	Width, Height int
}

func (r *rectInstruction) Execute(d *display) {
	for y := 0; y < r.Height; y++ {
		for x := 0; x < r.Width; x++ {
			d.Pixels[y][x] = true
		}
	}
}

type rotateRowInstruction struct {
	Y, N int
}

func (r *rotateRowInstruction) Execute(d *display) {
	w := len(d.Pixels[0])
	buf := make([]bool, w)

	// save current pixels to buffer
	for x := 0; x < w; x++ {
		buf[x] = d.Pixels[r.Y][x]
	}

	// save rotated pixels to display
	for x := 0; x < w; x++ {
		d.Pixels[r.Y][(x+r.N)%w] = buf[x]
	}
}

type rotateColumnInstruction struct {
	X, N int
}

func (r *rotateColumnInstruction) Execute(d *display) {
	h := len(d.Pixels)
	buf := make([]bool, h)

	// save current pixels to buffer
	for y := 0; y < h; y++ {
		buf[y] = d.Pixels[y][r.X]
	}

	// save rotated pixels to display
	for y := 0; y < h; y++ {
		d.Pixels[(y+r.N)%h][r.X] = buf[y]
	}
}
