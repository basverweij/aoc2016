// Code generated from .\Scramble.g4 by ANTLR 4.7.1. DO NOT EDIT.

package scramble // Scramble
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseScrambleListener is a complete listener for a parse tree produced by ScrambleParser.
type BaseScrambleListener struct{}

var _ ScrambleListener = &BaseScrambleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseScrambleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseScrambleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseScrambleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseScrambleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScrambler is called when production scrambler is entered.
func (s *BaseScrambleListener) EnterScrambler(ctx *ScramblerContext) {}

// ExitScrambler is called when production scrambler is exited.
func (s *BaseScrambleListener) ExitScrambler(ctx *ScramblerContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseScrambleListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseScrambleListener) ExitOperation(ctx *OperationContext) {}

// EnterSwapPosition is called when production swapPosition is entered.
func (s *BaseScrambleListener) EnterSwapPosition(ctx *SwapPositionContext) {}

// ExitSwapPosition is called when production swapPosition is exited.
func (s *BaseScrambleListener) ExitSwapPosition(ctx *SwapPositionContext) {}

// EnterSwapLetter is called when production swapLetter is entered.
func (s *BaseScrambleListener) EnterSwapLetter(ctx *SwapLetterContext) {}

// ExitSwapLetter is called when production swapLetter is exited.
func (s *BaseScrambleListener) ExitSwapLetter(ctx *SwapLetterContext) {}

// EnterRotateLeft is called when production rotateLeft is entered.
func (s *BaseScrambleListener) EnterRotateLeft(ctx *RotateLeftContext) {}

// ExitRotateLeft is called when production rotateLeft is exited.
func (s *BaseScrambleListener) ExitRotateLeft(ctx *RotateLeftContext) {}

// EnterRotateRight is called when production rotateRight is entered.
func (s *BaseScrambleListener) EnterRotateRight(ctx *RotateRightContext) {}

// ExitRotateRight is called when production rotateRight is exited.
func (s *BaseScrambleListener) ExitRotateRight(ctx *RotateRightContext) {}

// EnterRotateLetter is called when production rotateLetter is entered.
func (s *BaseScrambleListener) EnterRotateLetter(ctx *RotateLetterContext) {}

// ExitRotateLetter is called when production rotateLetter is exited.
func (s *BaseScrambleListener) ExitRotateLetter(ctx *RotateLetterContext) {}

// EnterReverse is called when production reverse is entered.
func (s *BaseScrambleListener) EnterReverse(ctx *ReverseContext) {}

// ExitReverse is called when production reverse is exited.
func (s *BaseScrambleListener) ExitReverse(ctx *ReverseContext) {}

// EnterMove is called when production move is entered.
func (s *BaseScrambleListener) EnterMove(ctx *MoveContext) {}

// ExitMove is called when production move is exited.
func (s *BaseScrambleListener) ExitMove(ctx *MoveContext) {}
