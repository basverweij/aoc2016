package main

import (
	"fmt"
	"strconv"
	"strings"
)

type nodeMap map[string]*node

type node struct {
	name                   string
	size, used, avail, use int
}

func (n *node) pos() (x, y int) {
	s := strings.Split(n.name, "-")

	x, _ = strconv.Atoi(s[1][1:])
	y, _ = strconv.Atoi(s[2][1:])

	return
}

func (m nodeMap) nodeAtPos(x, y int) *node {
	return m[fmt.Sprintf("/dev/grid/node-x%d-y%d", x, y)]
}

var adjacentMoves = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func (m nodeMap) adjacentNodes(n *node) []*node {
	adjs := make([]*node, 0)
	x, y := n.pos()

	for _, move := range adjacentMoves {
		if adj := m.nodeAtPos(x+move[0], y+move[1]); adj != nil {
			adjs = append(adjs, adj)
		}
	}

	return adjs
}
