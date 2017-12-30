package main

type context struct {
	left, center, right tile
}

var rules = map[context]tile{
	{trap, trap, safe}: trap,
	{safe, trap, trap}: trap,
	{trap, safe, safe}: trap,
	{safe, safe, trap}: trap,
}

func applyRules(row []tile) []tile {
	n := len(row)
	next := make([]tile, n)

	next[0] = rules[context{safe, row[0], row[1]}]

	for i := 1; i < n-1; i++ {
		next[i] = rules[context{row[i-1], row[i], row[i+1]}]
	}

	next[n-1] = rules[context{row[n-2], row[n-1], safe}]

	return next
}
