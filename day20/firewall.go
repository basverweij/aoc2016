package main

import (
	"math"
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

func findAmountAllowed(rules []rule, highest uint32) uint32 {
	sort.Sort(byFrom(rules))

	var amount uint32
	var ip uint32
	for i := 0; i < len(rules); i++ {
		if ip < rules[i].from {
			amount += rules[i].from - ip
		}

		if rules[i].to == math.MaxUint32 {
			return amount
		}

		ip = max(ip, rules[i].to+1)
	}

	amount += highest - ip + 1

	return amount
}
