package main

import "testing"

func TestDisplay(t *testing.T) {
	d := newDisplay(7, 3)

	lines := []string{
		"rect 3x2",
		"rotate column x=1 by 1",
		"rotate row y=0 by 4",
		"rotate column x=1 by 1",
	}

	for _, line := range lines {
		i, err := parse(line)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		i.Execute(d)

		t.Logf("\n%s", d.String())
	}
}
