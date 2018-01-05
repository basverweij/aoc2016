// Code generated from .\Scramble.g4 by ANTLR 4.7.1. DO NOT EDIT.

package scramble // Scramble
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ScrambleListener is a complete listener for a parse tree produced by ScrambleParser.
type ScrambleListener interface {
	antlr.ParseTreeListener

	// EnterScrambler is called when entering the scrambler production.
	EnterScrambler(c *ScramblerContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterSwapPosition is called when entering the swapPosition production.
	EnterSwapPosition(c *SwapPositionContext)

	// EnterSwapLetter is called when entering the swapLetter production.
	EnterSwapLetter(c *SwapLetterContext)

	// EnterRotateLeft is called when entering the rotateLeft production.
	EnterRotateLeft(c *RotateLeftContext)

	// EnterRotateRight is called when entering the rotateRight production.
	EnterRotateRight(c *RotateRightContext)

	// EnterRotateLetter is called when entering the rotateLetter production.
	EnterRotateLetter(c *RotateLetterContext)

	// EnterReverse is called when entering the reverse production.
	EnterReverse(c *ReverseContext)

	// EnterMove is called when entering the move production.
	EnterMove(c *MoveContext)

	// ExitScrambler is called when exiting the scrambler production.
	ExitScrambler(c *ScramblerContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitSwapPosition is called when exiting the swapPosition production.
	ExitSwapPosition(c *SwapPositionContext)

	// ExitSwapLetter is called when exiting the swapLetter production.
	ExitSwapLetter(c *SwapLetterContext)

	// ExitRotateLeft is called when exiting the rotateLeft production.
	ExitRotateLeft(c *RotateLeftContext)

	// ExitRotateRight is called when exiting the rotateRight production.
	ExitRotateRight(c *RotateRightContext)

	// ExitRotateLetter is called when exiting the rotateLetter production.
	ExitRotateLetter(c *RotateLetterContext)

	// ExitReverse is called when exiting the reverse production.
	ExitReverse(c *ReverseContext)

	// ExitMove is called when exiting the move production.
	ExitMove(c *MoveContext)
}
