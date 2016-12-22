package main

import "fmt"

type step struct {
	model   *model
	ToFloor uint32
	Devices []uint32
}

func newStep(m *model, toFloor uint32, devices ...uint32) step {
	return step{m, toFloor, devices}
}

func (s step) String() string {
	var codes []string
	for _, idx := range s.Devices {
		codes = append(codes, s.model.Code(idx))
	}

	return fmt.Sprintf("Move %v to floor %d", codes, s.ToFloor)
}

func (m *model) NextSteps() []step {
	var steps []step

	// get devices on current floor
	availableDevs := m.availableDevices()

	var availableFloors []uint32
	if m.Elevator() < 4 {
		// if we can move up, check this first
		availableFloors = append(availableFloors, m.Elevator()+1)
	}

	if m.Elevator() > 1 {
		// check if anything is left on the lower flowers
		var numLeft uint32
		for i := uint32(1); i < m.Elevator(); i++ {
			numLeft += uint32(len(m.devicesOnFloor(i)))
		}

		if numLeft > 0 {
			availableFloors = append(availableFloors, m.Elevator()-1)
		}
	}

	for _, toFloor := range availableFloors {
		for i, idx1 := range availableDevs {
			// all moves of two devices are a step
			for _, idx2 := range availableDevs[i+1:] {
				steps = append(steps, newStep(m, toFloor, idx1, idx2))
			}

			// all single device moves are a step
			steps = append(steps, newStep(m, toFloor, idx1))
		}

	}

	return steps
}

func (m *model) devicesOnFloor(floor uint32) []uint32 {
	var devs []uint32

	var i uint32
	for i = 0; i < m.NumDevices(); i++ {
		if m.getFloor(i) == floor {
			devs = append(devs, i)
		}
	}

	return devs
}

func (m *model) availableDevices() []uint32 {
	return m.devicesOnFloor(m.Elevator())
}
