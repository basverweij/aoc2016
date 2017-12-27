package main

func gen(dataLen int, startValue int64, startLen int) string {
	d := newDragon(startValue, startLen)
	for ; d.len < dataLen; d.expand() {
	}

	return checksum(d, dataLen)
}
