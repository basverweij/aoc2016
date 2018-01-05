package main

type pair struct {
	a, b *node
}

func viablePairs(nodes nodeMap) []pair {
	pairs := make([]pair, 0)

	for _, a := range nodes {
		if a.used == 0 {
			continue
		}

		for _, b := range nodes {
			if a == b {
				continue
			}

			if a.used <= b.avail {
				pairs = append(pairs, pair{a, b})
			}
		}
	}

	return pairs
}

func connectedPairs(nodes nodeMap) []pair {
	pairs := make([]pair, 0)

	for _, n := range nodes {
		if n.used == 0 {
			continue
		}

		for _, a := range nodes.adjacentNodes(n) {
			if n.used <= a.avail {
				pairs = append(pairs, pair{n, a})
			}
		}
	}

	return pairs
}
