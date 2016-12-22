package main

import "fmt"

var input *model
var targetHash uint64

func init() {
	input = newModel()
	input.AddDevice("Pr", 1, 1)
	input.AddDevice("Co", 2, 3)
	input.AddDevice("Cu", 2, 3)
	input.AddDevice("Ru", 2, 3)
	input.AddDevice("Pl", 2, 3)
	input.AddDevice("El", 1, 1)
	input.AddDevice("Di", 1, 1)

	target := newModel()
	target.MoveElevator(2)
	target.MoveElevator(3)
	target.MoveElevator(4)
	target.AddDevice("Pr", 4, 4)
	target.AddDevice("Co", 4, 4)
	target.AddDevice("Cu", 4, 4)
	target.AddDevice("Ru", 4, 4)
	target.AddDevice("Pl", 4, 4)
	target.AddDevice("El", 4, 4)
	target.AddDevice("Di", 4, 4)

	targetHash = target.Hash()
	fmt.Printf("target hash = %d (%032b)\n", targetHash, targetHash)
}

func initRef() {
	input = newModel()
	input.AddDevice("Hy", 2, 1)
	input.AddDevice("Li", 3, 1)

	target := newModel()
	target.MoveElevator(2)
	target.MoveElevator(3)
	target.MoveElevator(4)
	target.AddDevice("Hy", 4, 4)
	target.AddDevice("Li", 4, 4)

	targetHash = target.Hash()
	fmt.Printf("target hash = %d (%032b)\n", targetHash, targetHash)
}
