package main

import "fmt"

type system struct {
	instructions []instruction
	Reg          map[string]int
	PC           int
}

func newSystem(regs ...string) *system {
	sys := &system{Reg: make(map[string]int)}

	for _, reg := range regs {
		sys.Reg[reg] = 0
	}

	return sys
}

func (s *system) addInstruction(i instruction) {
	s.instructions = append(s.instructions, i)
}

func (s *system) Cycle() bool {
	s.instructions[s.PC].Execute(s)
	s.PC++

	return s.PC >= len(s.instructions)
}

func (s *system) Run() {
	for {
		if s.Cycle() {
			// stop if cycle returns true
			return
		}
	}
}

func (s *system) Print() {
	fmt.Printf("%8v @ %03d: %v\n", s.Reg, s.PC, s.instructions[s.PC])
}

type instruction interface {
	Execute(sys *system)
}
