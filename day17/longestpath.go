package main

func longestPath(passcode string) int {
	var longest *path

	q := new(pathQueue)
	for p := new(path); p != nil; p = q.pop() {
		dirs := getDirs(p, passcode)

		for _, dir := range dirs {
			e := p.extend(dir)

			if e.x == 3 && e.y == 3 {
				// we're at the vault
				if longest == nil || len(longest.moves) < len(e.moves) {
					longest = e
				}
			} else {
				q.push(e)
			}
		}
	}

	if longest == nil {
		return 0
	}

	return len(longest.moves)
}
