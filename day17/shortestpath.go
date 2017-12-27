package main

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
