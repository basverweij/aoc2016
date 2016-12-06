package main

import "testing"

func testGetMax(t *testing.T, freqs []int, expectedMax rune) {
	actualMax := getMax(freqs)
	if actualMax != expectedMax {
		t.Errorf("expected '%v' as max for %v, but got '%s'", expectedMax, freqs, actualMax)
	}
}

func TestGetMax(t *testing.T) {
	testGetMax(t, []int{0}, 'a')
	testGetMax(t, []int{1}, 'a')
	testGetMax(t, []int{0, 1}, 'b')
}

func testDecodeMsg(t *testing.T, msgs []string, expectedMsg string) {
	actualMsg := decodeMsg(msgs)
	if actualMsg != expectedMsg {
		t.Errorf("expected msg '%s', but got '%s'", expectedMsg, actualMsg)
	}
}

func TestDecodeMsg(t *testing.T) {
	msgs := []string{
		"eedadn",
		"drvtee",
		"eandsr",
		"raavrd",
		"atevrs",
		"tsrnev",
		"sdttsa",
		"rasrtv",
		"nssdts",
		"ntnada",
		"svetve",
		"tesnvt",
		"vntsnd",
		"vrdear",
		"dvrsen",
		"enarar",
	}

	testDecodeMsg(t, msgs, "easter")
}
