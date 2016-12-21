package main

import "fmt"

const elevatorIdx uint32 = 15

var deviceTypes = []string{"G", "M"}

type model struct {
	// codes contains the device codes
	codes []string

	// floors contains the floor of each device, encoded in 2 bits
	// starting with the first device in the least significant bits (0-1),
	// the second device in bits 2-3, etc.
	// the elevator position is saved in bits 30-31
	floors uint32
}

func newModel() *model {
	m := &model{}

	// set elevator to first floor
	m.setFloor(elevatorIdx, 1)

	return m
}

func (m *model) setFloor(idx, floor uint32) {
	// we use 2 bits per device
	idx <<= 1

	m.floors &= ^(uint32(3) << idx)
	m.floors |= (floor - 1) << idx
}

func (m *model) getFloor(idx uint32) uint32 {
	// we use 2 bits per device
	idx <<= 1

	return ((m.floors >> idx) & uint32(3)) + 1
}

func (m *model) AddDevice(code string, genFloor, chpFloor uint32) {
	// each device uses two pairs of 2 bits,
	// and the first 2 bits are for the elevator
	n := m.NumDevices()
	if n >= elevatorIdx-1 {
		panic("too many devices")
	}

	// add code
	m.codes = append(m.codes, code)

	// set generator and microchip floors
	m.setFloor(n, genFloor)
	m.setFloor(n+1, chpFloor)
}

func (m *model) Hash() uint32 {
	return m.floors
}

func (m *model) NumDevices() uint32 {
	return uint32(len(m.codes) * 2)
}

func (m *model) Code(idx uint32) string {
	return m.codes[idx>>1] + deviceTypes[idx&1]
}

func (m *model) Elevator() uint32 {
	return m.getFloor(elevatorIdx)
}

func (m *model) MoveElevator(toFloor uint32, devices ...uint32) {
	if toFloor != m.Elevator()-1 && toFloor != m.Elevator()+1 {
		panic("invalid elevator floor")
	}

	m.setFloor(elevatorIdx, toFloor)

	for _, device := range devices {
		m.setFloor(device, toFloor)
	}
}

func (m *model) Print() {
	var s, fs string

	var i, j uint32
	nCodes := uint32(len(m.codes))

	var f uint32 = 4
	for ; f >= 1; f-- {
		if m.Elevator() == f {
			s = "E"
		} else {
			s = "."
		}

		fs = fmt.Sprintf("F%d   %s    ", f, s)

		for i = 0; i < nCodes; i++ {
			for j = 0; j < 2; j++ {
				if m.getFloor(i*2+j) == f {
					s = m.codes[i] + deviceTypes[j]
				} else {
					s = ".  "
				}

				fs += fmt.Sprintf("%s  ", s)
			}
		}

		fmt.Println(fs)
	}
}
