package main

type step struct {
	Dev1, Dev2 string
	ToFloor    int
}

func (m *model) NextSteps() []step {
	var steps []step

	// get devices on current floor
	availableDevs := m.availableDevices()

	var availableFloors []int
	if m.Elevator > 1 {
		availableFloors = append(availableFloors, m.Elevator-1)
	}

	if m.Elevator < 4 {
		availableFloors = append(availableFloors, m.Elevator+1)
	}

	for _, toFloor := range availableFloors {
		for i, code := range availableDevs {
			// all single device moves are a step
			steps = append(steps, step{code, "", toFloor})

			// all moves of two devices are a step
			for _, code2 := range availableDevs[i+1:] {
				steps = append(steps, step{code, code2, toFloor})
			}
		}

	}

	return steps
}

func (m *model) availableDevices() []string {
	var devs []string

	for _, code := range m.devices {
		if m.floors[code] == m.Elevator {
			devs = append(devs, code)
		}
	}

	return devs
}
