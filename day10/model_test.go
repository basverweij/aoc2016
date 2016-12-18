package main

import "testing"

func TestEmptyModel(t *testing.T) {
	m := newModel(make([]*initInstruction, 0), make([]*compareInstruction, 0))

	b := m.step()
	assertNil(t, b, "step result")
}

func TestOnlyInitsModel(t *testing.T) {
	inits := make([]*initInstruction, 2)
	inits[0] = &initInstruction{BotID: 0, Value: 1}
	inits[1] = &initInstruction{BotID: 0, Value: 2}

	cmps := make([]*compareInstruction, 1)
	cmps[0] = &compareInstruction{BotID: 0, LowIsBot: false, Low: 1, HighIsBot: false, High: 0}

	m := newModel(inits, cmps)

	b := m.step()
	assertNotNil(t, b, "step result")
}

func TestModelSteps(t *testing.T) {
	inits := make([]*initInstruction, 3)
	inits[0] = &initInstruction{BotID: 2, Value: 5}
	inits[1] = &initInstruction{BotID: 1, Value: 3}
	inits[2] = &initInstruction{BotID: 2, Value: 2}

	cmps := make([]*compareInstruction, 3)
	cmps[0] = &compareInstruction{BotID: 2, LowIsBot: true, Low: 1, HighIsBot: true, High: 0}
	cmps[1] = &compareInstruction{BotID: 1, LowIsBot: false, Low: 1, HighIsBot: true, High: 0}
	cmps[2] = &compareInstruction{BotID: 0, LowIsBot: false, Low: 2, HighIsBot: false, High: 0}

	m := newModel(inits, cmps)

	bots := m.step()
	assertNotNil(t, bots, "step #1 result")
	assertEquals(t, 2, bots[0].BotID, "step #1 bot #1 id")
	assertEquals(t, 1, bots[1].BotID, "step #1 bot #2 id")
	assertEquals(t, 0, bots[2].BotID, "step #1 bot #3 id")
}
