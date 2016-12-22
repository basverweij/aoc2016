package main

import "fmt"

type cpyReg struct {
	From, To string
}

func (i *cpyReg) Execute(sys *system) {
	sys.Reg[i.To] = sys.Reg[i.From]
}

func (i *cpyReg) String() string {
	return fmt.Sprintf("cpy %s %s", i.From, i.To)
}

type cpyVal struct {
	Val int
	To  string
}

func (i *cpyVal) Execute(sys *system) {
	sys.Reg[i.To] = i.Val
}

func (i *cpyVal) String() string {
	return fmt.Sprintf("cpy %d %s", i.Val, i.To)
}

type inc struct {
	Reg string
}

func (i *inc) Execute(sys *system) {
	sys.Reg[i.Reg]++
}

func (i *inc) String() string {
	return fmt.Sprintf("inc %s", i.Reg)
}

type dec struct {
	Reg string
}

func (i *dec) Execute(sys *system) {
	sys.Reg[i.Reg]--
}

func (i *dec) String() string {
	return fmt.Sprintf("dec %s", i.Reg)
}

type jnzReg struct {
	Reg  string
	Jump int
}

func (i *jnzReg) Execute(sys *system) {
	if sys.Reg[i.Reg] != 0 {
		// we subtract 1 from jump because the system always moves the
		// program counter ahead by one
		sys.PC += i.Jump - 1
	}
}

func (i *jnzReg) String() string {
	return fmt.Sprintf("jnz %s %d", i.Reg, i.Jump)
}

type jnzVal struct {
	Val, Jump int
}

func (i *jnzVal) Execute(sys *system) {
	if i.Val != 0 {
		// we subtract 1 from jump because the system always moves the
		// program counter ahead by one
		sys.PC += i.Jump - 1
	}
}

func (i *jnzVal) String() string {
	return fmt.Sprintf("jnz %d %d", i.Val, i.Jump)
}
