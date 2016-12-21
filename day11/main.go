package main

import "fmt"

var n int

func main() {
	fmt.Println("\ninput:")
	input.Print()

	move("PrG", "PrM", 2)

	steps := input.NextSteps()
	for _, step := range steps {
		fmt.Printf("Move '%s' and '%s' to %d\n", step.Dev1, step.Dev2, step.ToFloor)
	}
}

func move(device1, device2 string, toFloor int) {
	n++
	fmt.Printf("\nstep #%d:\n", n)

	if err := input.Move(device1, device2, toFloor); err != nil {
		panic(err)
	}

	input.Print()
}
