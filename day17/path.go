package main

import "fmt"

type direction int

const (
	up direction = iota
	down
	left
	right
)

var (
	dirRunes = []rune{'U', 'D', 'L', 'R'}
	xMoves   = []int{0, 0, -1, 1}
	yMoves   = []int{-1, 1, 0, 0}
)

type path struct {
	moves []direction
	x, y  int
}

func (p *path) String() string {
	return fmt.Sprintf("@(%d, %d): %s", p.x, p.y, p.movesString())
}

func (p *path) movesString() string {
	s := make([]rune, len(p.moves))
	for i := range s {
		s[i] = dirRunes[p.moves[i]]
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
