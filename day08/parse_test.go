package main

import "testing"

func TestParseRectInstruction(t *testing.T) {
	i, err := parse("rect 1x2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r, isType := i.(*rectInstruction)
	if !isType {
		t.Errorf("expected type *main.rectInstruction, but got %T", i)
		return
	}

	if r.Width != 1 || r.Height != 2 {
		t.Errorf("expected dimensions 1x2 for rect instruction, but got %dx%d", r.Width, r.Height)
	}
}

func TestParseRotateRowInstruction(t *testing.T) {
	i, err := parse("rotate row y=1 by 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r, isType := i.(*rotateRowInstruction)
	if !isType {
		t.Errorf("expected type *main.rotateRowInstruction, but got %T", i)
		return
	}

	if r.Y != 1 || r.N != 2 {
		t.Errorf("expected Y=1 and N=2 for rotate row instruction, but got Y=%d and N=%d", r.Y, r.N)
	}
}

func TestParseRotateColumnInstruction(t *testing.T) {
	i, err := parse("rotate column x=1 by 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r, isType := i.(*rotateColumnInstruction)
	if !isType {
		t.Errorf("expected type *main.rotateColumnInstruction, but got %T", i)
		return
	}

	if r.X != 1 || r.N != 2 {
		t.Errorf("expected X=1 and N=2 for rotate row instruction, but got X=%d and N=%d", r.X, r.N)
	}
}
