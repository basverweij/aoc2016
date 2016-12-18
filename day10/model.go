package main

import "fmt"

type model struct {
	cmps map[int]*compareInstruction
	bots map[int]*bot

	// ValueCompares saves all low (first map index) and high (second map index) compares
	// and maps these to the last bot id that compared these two value
	ValueCompares map[int]map[int]int

	// Outputs stores the last value put in each output (map index)
	Outputs map[int]int
}

type bot struct {
	BotID int
	Low   int
	High  int
}

func newModel(inits []*initInstruction, cmps []*compareInstruction) *model {
	m := &model{}
	m.ValueCompares = make(map[int]map[int]int, 0)
	m.Outputs = make(map[int]int, 0)

	m.bots = make(map[int]*bot)

	for _, init := range inits {
		m.checkBot(init.BotID).addValue(init.Value)
	}

	m.cmps = make(map[int]*compareInstruction, len(cmps))
	for _, cmp := range cmps {
		m.cmps[cmp.BotID] = cmp
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

func (m *model) SaveValueCompare(b *bot) {
	lm, found := m.ValueCompares[b.Low]
	if !found {
		lm = make(map[int]int, 1)
		m.ValueCompares[b.Low] = lm
	}

	lm[b.High] = b.BotID
}

func (b *bot) addValue(value int) {
	if value < 1 {
		panic(fmt.Sprintf("invalid value: %d", value))
	}

	if b.Low == 0 {
		b.Low = value
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

func (b *bot) removeValue(value int) {
	if value == b.Low {
		b.Low, b.High = b.High, 0
	} else if value == b.High {
		b.High = 0
	} else {
		panic(fmt.Sprintf("invalid value %d for bot %v", value, b))
	}
}

// step executes one step in the model, and returns the bot that executed this step.
// It returns nil if no step can be executed, which indicates that the model is done.
func (m *model) step() []*bot {
	var bots []*bot

	for _, bot := range m.bots {
		if bot.Low == 0 || bot.High == 0 {
			continue
		}

		cmp, found := m.cmps[bot.BotID]
		if !found {
			panic(fmt.Sprintf("no compare instruction for bot #%d", bot.BotID))
		}

		if cmp.LowIsBot {
			m.checkBot(cmp.Low).addValue(bot.Low)
		} else {
			m.Outputs[cmp.Low] = bot.Low
		}

		if cmp.HighIsBot {
			m.checkBot(cmp.High).addValue(bot.High)
		} else {
			m.Outputs[cmp.High] = bot.High
		}

		m.SaveValueCompare(bot)

		bot.Low, bot.High = 0, 0

		bots = append(bots, bot)
	}

	return bots
}
