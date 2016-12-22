package main

import "testing"

func TestModel(t *testing.T) {
	m := newSystem("a")
	m.addInstruction(&cpyVal{Val: 41, To: "a"})
	m.addInstruction(&inc{Reg: "a"})
	m.addInstruction(&inc{Reg: "a"})
	m.addInstruction(&dec{Reg: "a"})
	m.addInstruction(&jnzReg{Reg: "a", Jump: 2})
	m.addInstruction(&dec{Reg: "a"})

	m.Run()

	assertEquals(t, 42, m.Reg["a"], "value of register 'a'")
}
