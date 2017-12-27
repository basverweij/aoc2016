package main

import "crypto/md5"

type direction int

const (
	up direction = iota
	down
	left
	right
)

var (
	dirBytes = []byte{'U', 'D', 'L', 'R'}
	xMoves   = []int{0, 0, -1, 1}
	yMoves   = []int{-1, 1, 0, 0}
)

func getDirs(path *path, passcode string) []direction {
	h := md5.Sum([]byte(passcode + path.movesString()))

	var dirs []direction
	for i := 0; i < 4; i++ {
		if ((h[i/2] >> uint(4-(i%2)*4)) & 0xf) < 0xb {
			// locked
			continue
		}

		if layout[path.y][path.x][i] {
			dirs = append(dirs, direction(i))
		}
	}

	return dirs
}
