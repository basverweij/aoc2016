package main

/*
func (m *model) Move(device1, device2 string, toFloor uint32) error {
	if toFloor != m.Elevator()-1 && toFloor != m.Elevator()+1 {
		return fmt.Errorf("invalid toFloor %d (current floor = %d)", toFloor, m.Elevator)
	}

	if device1 == "" && device2 == "" {
		return fmt.Errorf("specify at least one device to move")
	}

	if device1 != "" {
		if err := m.moveDevice(device1, toFloor); err != nil {
			return err
		}
	}

	if device2 != "" {
		if err := m.moveDevice(device2, toFloor); err != nil {
			return err
		}
	}

	m.setFloor(0, toFloor)

	return nil
}

func (m *model) moveDevice(code string, toFloor int) error {
	// check if device is currently on the same floor as the elevator
	if m.floors[code] != m.Elevator {
		return fmt.Errorf("%s is not on floor %d", code, m.Elevator)
	}

	m.floors[code] = toFloor

	return nil
}
*/
