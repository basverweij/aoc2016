package main

import (
	"fmt"
)

type model struct {
	cmps map[int]*compareInstruction
	bots map[int]*bot
}

type bot struct {
	BotID int
	Low   int
	High  int
}

func newModel(inits []*initInstruction, cmps []*compareInstruction) *model {
	m := &model{}

	m.bots = make(map[int]*bot)

	for _, init := range inits {
		m.checkBot(init.Bot).addValue(init.Value)
	}

	m.cmps = make(map[int]*compareInstruction, len(cmps))
	for _, cmp := range cmps {
		m.cmps[cmp.Bot] = cmp
		m.checkBot(cmp.Bot)
	}

	return m
}

func (m *model) checkBot(botID int) *bot {
	if bot, found := m.bots[botID]; found {
		return bot
	}

	bot := &bot{BotID: botID}
	m.bots[botID] = bot

	return bot
}

func (b *bot) addValue(value int) {
	if value < 1 {
		panic(fmt.Sprintf("invalid value: %d", value))
	}

	if b.Low == 0 {
		if value < b.High {
			b.Low = value
		} else {
			b.Low, b.High = b.High, value
		}
	} else if b.High == 0 {
		if value < b.Low {
			b.Low, b.High = value, b.Low
		} else {
			b.High = value
		}
	} else {
		panic("both low and high values are already set")
	}
}

// step executes one step in the model, and returns the bot that executed this step.
// It returns nil if no step can be executed, which indicates that the model is done.
func (m *model) step() *bot {
	for _, bot := range m.bots {
		if bot.Low == 0 || bot.High == 0 {
			continue
		}

		return bot
	}

	return nil
}
