package main

import "testing"

func TestParseEmptyLine(t *testing.T) {
	_, _, err := parse("")
	assertNotEquals(t, nil, err, "error")
}

func TestParseInitInstruction(t *testing.T) {
	init, _, err := parse("value 5 goes to bot 2")
	assertEquals(t, nil, err, "error")
	assertNotEquals(t, nil, init, "init instruction")
	assertEquals(t, 5, init.Value, "value")
	assertEquals(t, 2, init.BotID, "bot")
}

func TestParseCompareInstructionWithOnlyBots(t *testing.T) {
	_, cmp, err := parse("bot 2 gives low to bot 1 and high to bot 0")
	assertEquals(t, nil, err, "error")
	assertNotEquals(t, nil, cmp, "compare instruction")
	assertEquals(t, 2, cmp.BotID, "bot")
	assertEquals(t, true, cmp.LowIsBot, "low is bot")
	assertEquals(t, 1, cmp.Low, "low")
	assertEquals(t, true, cmp.LowIsBot, "high is bot")
	assertEquals(t, 0, cmp.High, "high")
}

func TestParseCompareInstructionWithBotAndOutput(t *testing.T) {
	_, cmp, err := parse("bot 2 gives low to bot 1 and high to output 0")
	assertEquals(t, nil, err, "error")
	assertNotEquals(t, nil, cmp, "compare instruction")
	assertEquals(t, 2, cmp.BotID, "bot")
	assertEquals(t, true, cmp.LowIsBot, "low is bot")
	assertEquals(t, 1, cmp.Low, "low")
	assertEquals(t, false, cmp.HighIsBot, "high is bot")
	assertEquals(t, 0, cmp.High, "high")
}

func TestParseCompareInstructionWithOutputAndBot(t *testing.T) {
	_, cmp, err := parse("bot 2 gives low to output 1 and high to bot 0")
	assertEquals(t, nil, err, "error")
	assertNotEquals(t, nil, cmp, "compare instruction")
	assertEquals(t, 2, cmp.BotID, "bot")
	assertEquals(t, false, cmp.LowIsBot, "low is bot")
	assertEquals(t, 1, cmp.Low, "low")
	assertEquals(t, true, cmp.HighIsBot, "high is bot")
	assertEquals(t, 0, cmp.High, "high")
}

func TestParseCompareInstructionWithOnlyOutputs(t *testing.T) {
	_, cmp, err := parse("bot 2 gives low to output 1 and high to output 0")
	assertEquals(t, nil, err, "error")
	assertNotEquals(t, nil, cmp, "compare instruction")
	assertEquals(t, 2, cmp.BotID, "bot")
	assertEquals(t, false, cmp.LowIsBot, "low is bot")
	assertEquals(t, 1, cmp.Low, "low")
	assertEquals(t, false, cmp.HighIsBot, "high is bot")
	assertEquals(t, 0, cmp.High, "high")
}
