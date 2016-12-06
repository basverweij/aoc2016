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
		msg[i] = getMax(freqs[i])
	}

	return string(msg)
}

func processMsg(freqs [][]int, msg string) {
	for i, r := range msg {
		freqs[i][r-'a']++
	}
}

func getMax(freqs []int) rune {
	max := 0
	c := 0

	for i, f := range freqs {
		if f > max {
			max = f
			c = i
		}
	}

	return rune('a' + c)
}
