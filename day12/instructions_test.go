package main

import "testing"

func getSystem() *system {
	s := newSystem("a", "b", "z")
	s.Reg["a"] = 1
	s.Reg["b"] = 2
	s.PC = 3

	return s
}

func checkInstruction(i instruction, check func(*system)) {
	s := getSystem()
	i.Execute(s)
	check(s)
}

func TestCpyVal(t *testing.T) {
	checkInstruction(&cpyVal{Val: 1, To: "a"}, func(s *system) {
		assertEquals(t, 1, s.Reg["a"], "value in register 'a'")
	})
}

func TestCpyReg(t *testing.T) {
	checkInstruction(&cpyReg{From: "b", To: "a"}, func(s *system) {
		assertEquals(t, 2, s.Reg["a"], "value in register 'a'")
	})
}

func TestInc(t *testing.T) {
	checkInstruction(&inc{Reg: "a"}, func(s *system) {
		assertEquals(t, 2, s.Reg["a"], "value in register 'a'")
	})

	checkInstruction(&inc{Reg: "b"}, func(s *system) {
		assertEquals(t, 3, s.Reg["b"], "value in register 'a'")
	})
}

func TestDec(t *testing.T) {
	checkInstruction(&dec{Reg: "a"}, func(s *system) {
		assertEquals(t, 0, s.Reg["a"], "value in register 'a'")
	})

	checkInstruction(&dec{Reg: "b"}, func(s *system) {
		assertEquals(t, 1, s.Reg["b"], "value in register 'a'")
	})
}

func TestJnzRegZeroReg(t *testing.T) {
	checkInstruction(&jnzReg{Reg: "z", Jump: 1}, func(s *system) {
		assertEquals(t, 3, s.PC, "program counter")
	})

	checkInstruction(&jnzReg{Reg: "z", Jump: 2}, func(s *system) {
		assertEquals(t, 3, s.PC, "program counter")
	})

	checkInstruction(&jnzReg{Reg: "z", Jump: 0}, func(s *system) {
		assertEquals(t, 3, s.PC, "program counter")
	})

	checkInstruction(&jnzReg{Reg: "z", Jump: -1}, func(s *system) {
		assertEquals(t, 3, s.PC, "program counter")
	})
}

func TestJnzReg(t *testing.T) {
	checkInstruction(&jnzReg{Reg: "a", Jump: 1}, func(s *system) {
		assertEquals(t, 3, s.PC, "program counter")
	})

	checkInstruction(&jnzReg{Reg: "a", Jump: 2}, func(s *system) {
		assertEquals(t, 4, s.PC, "program counter")
	})

	checkInstruction(&jnzReg{Reg: "a", Jump: 0}, func(s *system) {
		assertEquals(t, 2, s.PC, "program counter")
	})

	checkInstruction(&jnzReg{Reg: "a", Jump: -1}, func(s *system) {
		assertEquals(t, 1, s.PC, "program counter")
	})
}
