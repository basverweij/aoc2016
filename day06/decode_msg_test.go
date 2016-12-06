package main

import "testing"

func testGetMax(t *testing.T, freqs []int, expectedMax rune) {
	actualMax := getMax(freqs)
	if actualMax != expectedMax {
		t.Errorf("expected '%v' as max for %v, but got '%v'", expectedMax, freqs, actualMax)
	}
}

func TestGetMax(t *testing.T) {
	testGetMax(t, []int{1}, 'a')
	testGetMax(t, []int{2}, 'a')
	testGetMax(t, []int{1, 2}, 'b')
}

func testGetMin(t *testing.T, freqs []int, expectedMin rune) {
	actualMin := getMin(freqs)
	if actualMin != expectedMin {
		t.Errorf("expected '%v' as max for %v, but got '%v'", expectedMin, freqs, actualMin)
	}
}

func TestGetMin(t *testing.T) {
	testGetMin(t, []int{1}, 'a')
	testGetMin(t, []int{2}, 'a')
	testGetMin(t, []int{2, 1}, 'b')
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

	testDecodeMsg(t, msgs, "advent")
}
