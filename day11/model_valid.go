package main

func (m *model) IsValid() bool {
	// only check microchips
	for i := uint32(1); i < m.NumDevices(); i += 2 {
		if m.getFloor(i) == m.getFloor(i-1) {
			// microchip on same floor as its compatible generator
			continue
		}

		for j := uint32(0); j < m.NumDevices(); j += 2 {
			if j == i-1 {
				// skip the compatible generator
				continue
			}

			if m.getFloor(i) == m.getFloor(j) {
				// non-compatible generator on the same floor
				// (and compatible generator not found earlier)
				return false
			}
		}
	}

	// all checks passed
	return true
}
