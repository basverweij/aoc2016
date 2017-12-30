package main

import (
	"sort"
)

func findLowestAllowed(rules []rule) uint32 {
	sort.Sort(byFrom(rules))

	var ip uint32
	for i := 0; ip >= rules[i].from; i++ {
		ip = max(ip, rules[i].to+1)
	}

	return ip
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}

	return b
}
