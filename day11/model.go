package main

import "fmt"
import "sort"

type model struct {
	// Devices contains a sorted listed of all devices
	devices []string

	// Floors maps each device code (XxG or XxM) to its current floor
	floors map[string]int

	// Elevator contains the elevator's current floor
	Elevator int
}

func newModel() *model {
	return &model{floors: make(map[string]int)}
}

func (m *model) AddDevice(code string, floor int) {
	m.devices = append(m.devices, code)
	sort.Strings(m.devices)

	m.floors[code] = floor
}

func (m *model) Print() {
	var s, fs string
	for f := 4; f >= 1; f-- {
		s = "."
		if m.Elevator == f {
			s = "E"
		}

		fs = fmt.Sprintf("F%d   %s    ", f, s)

		for _, code := range m.devices {
			s = ".  "
			if m.floors[code] == f {
				s = code
			}

			fs += fmt.Sprintf("%s  ", s)
		}

		fmt.Println(fs)
	}
}
