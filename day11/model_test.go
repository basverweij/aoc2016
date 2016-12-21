package main

import "testing"

func TestSetElevator(t *testing.T) {
	m := newModel()
	assertEquals(t, uint32(0<<(elevatorIdx*2)), m.floors, "elevator position in floors")
	assertEquals(t, uint32(1), m.Elevator(), "elevator position")

	m.MoveElevator(2)
	assertEquals(t, uint32(1<<(elevatorIdx*2)), m.floors, "elevator position in floors")
	assertEquals(t, uint32(2), m.Elevator(), "elevator position")

	m.MoveElevator(3)
	assertEquals(t, uint32(2<<(elevatorIdx*2)), m.floors, "elevator position in floors")
	assertEquals(t, uint32(3), m.Elevator(), "elevator position")

	m.MoveElevator(4)
	assertEquals(t, uint32(3<<(elevatorIdx*2)), m.floors, "elevator position in floors")
	assertEquals(t, uint32(4), m.Elevator(), "elevator position")
}

func TestAddDevice(t *testing.T) {
	m := newModel()
	assertEquals(t, uint32(0), m.NumDevices(), "number of devices")

	m.MoveElevator(2)
	m.AddDevice("Te", 2, 3)
	assertEquals(t, uint32(2), m.NumDevices(), "number of devices")
	assertEquals(t, uint32((1<<(elevatorIdx*2))|(2<<2)|1), m.floors, "floors")
}
