package main

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/basverweij/aoc2016/day21/scramble"
)

func input(s string) []operation {
	input := antlr.NewInputStream(s)
	lexer := scramble.NewScrambleLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := scramble.NewScrambleParser(stream)
	p.BuildParseTrees = true

	blueprint := p.Scrambler()
	listener := &scrambleListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, blueprint)

	return listener.ops
}

type scrambleListener struct {
	*scramble.BaseScrambleListener

	ops []operation
	op  operation
}

func num(node antlr.TerminalNode) int {
	n, _ := strconv.Atoi(node.GetText())
	return n
}

func letter(node antlr.TerminalNode) byte {
	return node.GetText()[0]
}

func (s *scrambleListener) ExitOperation(ctx *scramble.OperationContext) {
	s.ops = append(s.ops, s.op)
	s.op = nil
}

func (s *scrambleListener) EnterSwapPosition(ctx *scramble.SwapPositionContext) {
	s.op = &swapPosition{
		from: num(ctx.Num(0)),
		to:   num(ctx.Num(1)),
	}
}

func (s *scrambleListener) EnterSwapLetter(ctx *scramble.SwapLetterContext) {
	s.op = &swapLetter{
		fromLetter: letter(ctx.Letter(0)),
		toLetter:   letter(ctx.Letter(1)),
	}
}

func (s *scrambleListener) EnterRotateLeft(ctx *scramble.RotateLeftContext) {
	s.op = &rotate{
		rotateLeft: true,
		steps:      num(ctx.Num()),
	}
}

func (s *scrambleListener) EnterRotateRight(ctx *scramble.RotateRightContext) {
	s.op = &rotate{
		rotateLeft: false,
		steps:      num(ctx.Num()),
	}
}

func (s *scrambleListener) EnterRotateLetter(ctx *scramble.RotateLetterContext) {
	s.op = &rotateLetter{
		letter: letter(ctx.Letter()),
	}
}

func (s *scrambleListener) EnterReverse(ctx *scramble.ReverseContext) {
	s.op = &reverse{
		from: num(ctx.Num(0)),
		to:   num(ctx.Num(1)),
	}
}

func (s *scrambleListener) EnterMove(ctx *scramble.MoveContext) {
	s.op = &move{
		from: num(ctx.Num(0)),
		to:   num(ctx.Num(1)),
	}
}
