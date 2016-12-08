package main

import "testing"

func TestRect(t *testing.T) {
	d := newDisplay(3, 3)
	i := rectInstruction{Width: 2, Height: 2}

	i.Execute(d)

	compareRow(t, d, 0, []bool{true, true, false})
	compareRow(t, d, 1, []bool{true, true, false})
	compareRow(t, d, 2, []bool{false, false, false})
}

func compareRow(t *testing.T, d *display, y int, expected []bool) {
	w := len(d.Pixels[y])

	if len(expected) != w {
		t.Errorf("row %d has width %d, but expected row has width %d", y, w, len(expected))
		return
	}

	for x := 0; x < w; x++ {
		if d.Pixels[y][x] != expected[x] {
			t.Errorf("expected %v at row %d column %d, but got %v", expected[x], y, x, d.Pixels[y][x])
		}
	}
}

func TestRotateRow(t *testing.T) {
	d := newDisplay(3, 1)
	i := &rotateRowInstruction{Y: 0, N: 2}

	d.Pixels[0][0] = true
	d.Pixels[0][1] = true

	i.Execute(d)
	compareRow(t, d, 0, []bool{true, false, true})

	i.Execute(d)
	compareRow(t, d, 0, []bool{false, true, true})

	i.Execute(d)
	compareRow(t, d, 0, []bool{true, true, false})
}

func compareColumn(t *testing.T, d *display, x int, expected []bool) {
	h := len(d.Pixels)

	if len(expected) != h {
		t.Errorf("column %d has height %d, but expected column has height %d", x, h, len(expected))
		return
	}

	for y := 0; y < h; y++ {
		if d.Pixels[y][x] != expected[y] {
			t.Errorf("expected %v at row %d column %d, but got %v", expected[y], y, x, d.Pixels[y][x])
		}
	}
}

func TestRotateColumn(t *testing.T) {
	d := newDisplay(1, 3)
	i := &rotateColumnInstruction{X: 0, N: 2}

	d.Pixels[0][0] = true
	d.Pixels[1][0] = true

	i.Execute(d)
	compareColumn(t, d, 0, []bool{true, false, true})

	i.Execute(d)
	compareColumn(t, d, 0, []bool{false, true, true})

	i.Execute(d)
	compareColumn(t, d, 0, []bool{true, true, false})
}
