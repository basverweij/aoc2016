package main

type pos struct {
	x, y int
}

type path struct {
	head pos
	len  int
}

func shortestPath(o office, from, to pos) int {
	todo := []path{path{head: from, len: 0}}

	seen := make(map[pos]struct{}, 0)

	var p path
	for len(todo) > 0 {
		p = todo[0]
		todo = todo[1:]

		if _, found := seen[p.head]; found {
			// already visited
			continue
		}

		seen[p.head] = struct{}{}

		if p.head == to {
			// found destination
			return p.len
		}

		if p.head.y > 0 && !o[p.head.y-1][p.head.x] {
			// top is open
			todo = append(todo, path{head: pos{p.head.x, p.head.y - 1}, len: p.len + 1})
		}

		if p.head.y < len(o)-1 && !o[p.head.y+1][p.head.x] {
			// bottom is open
			todo = append(todo, path{head: pos{p.head.x, p.head.y + 1}, len: p.len + 1})
		}

		if p.head.x > 0 && !o[p.head.y][p.head.x-1] {
			// left is open
			todo = append(todo, path{head: pos{p.head.x - 1, p.head.y}, len: p.len + 1})
		}

		if p.head.x < len(o[0])-1 && !o[p.head.y][p.head.x+1] {
			// right is open
			todo = append(todo, path{head: pos{p.head.x + 1, p.head.y}, len: p.len + 1})
		}
	}

	return -1
}
