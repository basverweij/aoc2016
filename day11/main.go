package main

import "fmt"

// seen saves all previous seen model hashes,
// mapped against the step in which the hash was found.
var seen = make(map[uint32]int)

var minSteps = 1 << 31

func main() {
	fmt.Println("\ninput:")
	input.Print()

	fmt.Println(nextStep(*input, 0))

	fmt.Printf("minimum number of steps = %d\n", minSteps)
}

func nextStep(m model, numStep int) (bool, int) {
	if numStep >= minSteps {
		// already found a lower number of steps
		return false, numStep
	}

	h := m.Hash()

	if h == targetHash {
		return true, numStep
	}

	if seenSteps, found := seen[h]; found && seenSteps <= numStep {
		return false, numStep
	}

	seen[h] = numStep

	steps := m.NextSteps()
	// fmt.Printf("Step %03d: available steps = %v\n", numStep, steps)

	for _, step := range steps {
		n := &model{}
		*n = m

		n.MoveElevator(step.ToFloor, step.Devices...)

		// fmt.Printf("Step #%03d: %s (current hash = %032b)\n", numStep+1, step, n.Hash())
		// n.Print()

		if !n.IsValid() {
			// fmt.Println("not valid")
			continue
		}

		if found, foundSteps := nextStep(*n, numStep+1); found {
			fmt.Printf("found target after %d steps\n", foundSteps)

			if foundSteps < minSteps {
				minSteps = foundSteps
			}
		}
	}

	return false, numStep
}
