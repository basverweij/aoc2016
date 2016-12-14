package main

import "testing"

func TestEmptyModel(t *testing.T) {
	m := newModel(make([]*initInstruction, 0), make([]*compareInstruction, 0))

	b := m.step()
	assertNil(t, b, "step result")
}

func TestOnlyInitsModel(t *testing.T) {
	inits := make([]*initInstruction, 2)
	inits[0] = &initInstruction{Bot: 0, Value: 1}
	inits[1] = &initInstruction{Bot: 0, Value: 2}

	m := newModel(inits, make([]*compareInstruction, 0))

	b := m.step()
	assertNotNil(t, b, "step result")
}

func TestModelSteps(t *testing.T) {
	inits := make([]*initInstruction, 3)
	inits[0] = &initInstruction{Bot: 2, Value: 5}
	inits[1] = &initInstruction{Bot: 1, Value: 3}
	inits[2] = &initInstruction{Bot: 2, Value: 2}

	cmps := make([]*compareInstruction, 3)
	cmps[0] = &compareInstruction{Bot: 2, LowIsBot: true, Low: 1, HighIsBot: true, High: 0}
	cmps[1] = &compareInstruction{Bot: 1, LowIsBot: false, Low: 1, HighIsBot: true, High: 0}
	cmps[2] = &compareInstruction{Bot: 0, LowIsBot: false, Low: 2, HighIsBot: false, High: 0}

	m := newModel(inits, cmps)

	b := m.step()
	assertNotNil(t, b, "step #1 result")
	assertEquals(t, 2, b.BotID, "step #1 bot id")

	b = m.step()
	assertNotNil(t, b, "step #2 result")
	assertEquals(t, 1, b.BotID, "step #2 bot id")

	b = m.step()
	assertNotNil(t, b, "step #3 result")
	assertEquals(t, 0, b.BotID, "step #3 bot id")
}
