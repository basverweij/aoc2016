package main

import "fmt"

type path struct {
	moves []direction
	x, y  int
}

func (p *path) String() string {
	return fmt.Sprintf("@(%d, %d): %s", p.x, p.y, p.movesString())
}

func (p *path) movesString() string {
	s := make([]byte, len(p.moves))
	for i := range s {
		s[i] = dirBytes[p.moves[i]]
	}

	return string(s)
}

func (p *path) extend(dir direction) *path {
	moves := make([]direction, len(p.moves)+1)
	copy(moves, p.moves)
	moves[len(p.moves)] = dir

	return &path{
		moves: moves,
		x:     p.x + xMoves[dir],
		y:     p.y + yMoves[dir],
	}
}
