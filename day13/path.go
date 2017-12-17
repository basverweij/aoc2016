package main

type pos struct {
	x, y int
}

type path struct {
	head pos
	len  int
}

func shortestPath(o office, from pos, done func(path, map[pos]struct{}) int) int {
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

		d := done(p, seen)
		if d > 0 {
			return d
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

func findShortesPathTo(to pos) func(path, map[pos]struct{}) int {
	return func(p path, seen map[pos]struct{}) int {
		if p.head == to {
			return p.len
		}

		return 0
	}
}

func findUniqueLocationsAtLen(l int) func(path, map[pos]struct{}) int {
	return func(p path, seen map[pos]struct{}) int {
		if p.len > l {
			return len(seen) - 1
		}

		return 0
	}
}
