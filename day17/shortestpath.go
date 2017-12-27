package main

import (
	"crypto/md5"
)

func shortestPath(passcode string) string {
	var shortest *path

	q := new(pathQueue)
	for p := new(path); p != nil; p = q.pop() {
		dirs := getDirs(p, passcode)

		for _, dir := range dirs {
			e := p.extend(dir)

			if e.x == 3 && e.y == 3 {
				// we're at the vault
				if shortest == nil || len(shortest.moves) > len(e.moves) {
					shortest = e
				}
			} else {
				q.push(e)
			}
		}

		if shortest != nil && len(q.peek().moves) >= len(shortest.moves) {
			// there will never be a shorter path
			break
		}
	}

	if shortest == nil {
		return ""
	}

	return shortest.movesString()
}

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
