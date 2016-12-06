package main

func decodeMsg(msgs []string) string {
	msgLen := len(msgs[0])

	// init frequency table
	freqs := make([][]int, msgLen)
	for i := 0; i < msgLen; i++ {
		freqs[i] = make([]int, 26)
	}

	// loop msgs
	for _, msg := range msgs {
		processMsg(freqs, msg)
	}

	// build msg
	msg := make([]rune, msgLen)
	for i := 0; i < msgLen; i++ {
		msg[i] = getMin(freqs[i])
	}

	return string(msg)
}

func processMsg(freqs [][]int, msg string) {
	for i, r := range msg {
		freqs[i][r-'a']++
	}
}

func getMax(freqs []int) rune {
	max := freqs[0]
	c := 0

	for i, f := range freqs[1:] {
		if f > 0 && f > max {
			max = f
			c = i + 1
		}
	}

	return rune('a' + c)
}

func getMin(freqs []int) rune {
	min := freqs[0]
	c := 0

	for i, f := range freqs[1:] {
		if f > 0 && f < min {
			min = f
			c = i + 1
		}
	}

	return rune('a' + c)
}
