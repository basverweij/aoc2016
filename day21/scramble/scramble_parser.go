// Code generated from .\Scramble.g4 by ANTLR 4.7.1. DO NOT EDIT.

package scramble // Scramble
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 72, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 7, 2, 24, 10, 2, 12,
	2, 14, 2, 27, 11, 2, 3, 2, 5, 2, 30, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3, 3, 2, 8,
	9, 2, 70, 2, 20, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 40, 3, 2, 2, 2, 8, 45,
	3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 54, 3, 2, 2, 2, 14, 58, 3, 2, 2, 2,
	16, 61, 3, 2, 2, 2, 18, 66, 3, 2, 2, 2, 20, 25, 5, 4, 3, 2, 21, 22, 7,
	18, 2, 2, 22, 24, 5, 4, 3, 2, 23, 21, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25,
	23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2,
	2, 28, 30, 7, 18, 2, 2, 29, 28, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 3,
	3, 2, 2, 2, 31, 39, 5, 6, 4, 2, 32, 39, 5, 8, 5, 2, 33, 39, 5, 10, 6, 2,
	34, 39, 5, 12, 7, 2, 35, 39, 5, 14, 8, 2, 36, 39, 5, 16, 9, 2, 37, 39,
	5, 18, 10, 2, 38, 31, 3, 2, 2, 2, 38, 32, 3, 2, 2, 2, 38, 33, 3, 2, 2,
	2, 38, 34, 3, 2, 2, 2, 38, 35, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 37,
	3, 2, 2, 2, 39, 5, 3, 2, 2, 2, 40, 41, 7, 3, 2, 2, 41, 42, 7, 16, 2, 2,
	42, 43, 7, 4, 2, 2, 43, 44, 7, 16, 2, 2, 44, 7, 3, 2, 2, 2, 45, 46, 7,
	5, 2, 2, 46, 47, 7, 17, 2, 2, 47, 48, 7, 6, 2, 2, 48, 49, 7, 17, 2, 2,
	49, 9, 3, 2, 2, 2, 50, 51, 7, 7, 2, 2, 51, 52, 7, 16, 2, 2, 52, 53, 9,
	2, 2, 2, 53, 11, 3, 2, 2, 2, 54, 55, 7, 10, 2, 2, 55, 56, 7, 16, 2, 2,
	56, 57, 9, 2, 2, 2, 57, 13, 3, 2, 2, 2, 58, 59, 7, 11, 2, 2, 59, 60, 7,
	17, 2, 2, 60, 15, 3, 2, 2, 2, 61, 62, 7, 12, 2, 2, 62, 63, 7, 16, 2, 2,
	63, 64, 7, 13, 2, 2, 64, 65, 7, 16, 2, 2, 65, 17, 3, 2, 2, 2, 66, 67, 7,
	14, 2, 2, 67, 68, 7, 16, 2, 2, 68, 69, 7, 15, 2, 2, 69, 70, 7, 16, 2, 2,
	70, 19, 3, 2, 2, 2, 5, 25, 29, 38,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'swap position '", "' with position '", "'swap letter '", "' with letter '",
	"'rotate left '", "' steps'", "' step'", "'rotate right '", "'rotate based on position of letter '",
	"'reverse positions '", "' through '", "'move position '", "' to position '",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "Num", "Letter",
	"Newline",
}

var ruleNames = []string{
	"scrambler", "operation", "swapPosition", "swapLetter", "rotateLeft", "rotateRight",
	"rotateLetter", "reverse", "move",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ScrambleParser struct {
	*antlr.BaseParser
}

func NewScrambleParser(input antlr.TokenStream) *ScrambleParser {
	this := new(ScrambleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Scramble.g4"

	return this
}

// ScrambleParser tokens.
const (
	ScrambleParserEOF     = antlr.TokenEOF
	ScrambleParserT__0    = 1
	ScrambleParserT__1    = 2
	ScrambleParserT__2    = 3
	ScrambleParserT__3    = 4
	ScrambleParserT__4    = 5
	ScrambleParserT__5    = 6
	ScrambleParserT__6    = 7
	ScrambleParserT__7    = 8
	ScrambleParserT__8    = 9
	ScrambleParserT__9    = 10
	ScrambleParserT__10   = 11
	ScrambleParserT__11   = 12
	ScrambleParserT__12   = 13
	ScrambleParserNum     = 14
	ScrambleParserLetter  = 15
	ScrambleParserNewline = 16
)

// ScrambleParser rules.
const (
	ScrambleParserRULE_scrambler    = 0
	ScrambleParserRULE_operation    = 1
	ScrambleParserRULE_swapPosition = 2
	ScrambleParserRULE_swapLetter   = 3
	ScrambleParserRULE_rotateLeft   = 4
	ScrambleParserRULE_rotateRight  = 5
	ScrambleParserRULE_rotateLetter = 6
	ScrambleParserRULE_reverse      = 7
	ScrambleParserRULE_move         = 8
)

// IScramblerContext is an interface to support dynamic dispatch.
type IScramblerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScramblerContext differentiates from other interfaces.
	IsScramblerContext()
}

type ScramblerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScramblerContext() *ScramblerContext {
	var p = new(ScramblerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_scrambler
	return p
}

func (*ScramblerContext) IsScramblerContext() {}

func NewScramblerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScramblerContext {
	var p = new(ScramblerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_scrambler

	return p
}

func (s *ScramblerContext) GetParser() antlr.Parser { return s.parser }

func (s *ScramblerContext) AllOperation() []IOperationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationContext)(nil)).Elem())
	var tst = make([]IOperationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationContext)
		}
	}

	return tst
}

func (s *ScramblerContext) Operation(i int) IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *ScramblerContext) AllNewline() []antlr.TerminalNode {
	return s.GetTokens(ScrambleParserNewline)
}

func (s *ScramblerContext) Newline(i int) antlr.TerminalNode {
	return s.GetToken(ScrambleParserNewline, i)
}

func (s *ScramblerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScramblerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScramblerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterScrambler(s)
	}
}

func (s *ScramblerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitScrambler(s)
	}
}

func (p *ScrambleParser) Scrambler() (localctx IScramblerContext) {
	localctx = NewScramblerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ScrambleParserRULE_scrambler)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Operation()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(19)
				p.Match(ScrambleParserNewline)
			}
			{
				p.SetState(20)
				p.Operation()
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScrambleParserNewline {
		{
			p.SetState(26)
			p.Match(ScrambleParserNewline)
		}

	}

	return localctx
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_operation
	return p
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) SwapPosition() ISwapPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwapPositionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwapPositionContext)
}

func (s *OperationContext) SwapLetter() ISwapLetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwapLetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwapLetterContext)
}

func (s *OperationContext) RotateLeft() IRotateLeftContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRotateLeftContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRotateLeftContext)
}

func (s *OperationContext) RotateRight() IRotateRightContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRotateRightContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRotateRightContext)
}

func (s *OperationContext) RotateLetter() IRotateLetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRotateLetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRotateLetterContext)
}

func (s *OperationContext) Reverse() IReverseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReverseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReverseContext)
}

func (s *OperationContext) Move() IMoveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoveContext)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *ScrambleParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ScrambleParserRULE_operation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScrambleParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.SwapPosition()
		}

	case ScrambleParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.SwapLetter()
		}

	case ScrambleParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)
			p.RotateLeft()
		}

	case ScrambleParserT__7:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(32)
			p.RotateRight()
		}

	case ScrambleParserT__8:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(33)
			p.RotateLetter()
		}

	case ScrambleParserT__9:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(34)
			p.Reverse()
		}

	case ScrambleParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(35)
			p.Move()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISwapPositionContext is an interface to support dynamic dispatch.
type ISwapPositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwapPositionContext differentiates from other interfaces.
	IsSwapPositionContext()
}

type SwapPositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwapPositionContext() *SwapPositionContext {
	var p = new(SwapPositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_swapPosition
	return p
}

func (*SwapPositionContext) IsSwapPositionContext() {}

func NewSwapPositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwapPositionContext {
	var p = new(SwapPositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_swapPosition

	return p
}

func (s *SwapPositionContext) GetParser() antlr.Parser { return s.parser }

func (s *SwapPositionContext) AllNum() []antlr.TerminalNode {
	return s.GetTokens(ScrambleParserNum)
}

func (s *SwapPositionContext) Num(i int) antlr.TerminalNode {
	return s.GetToken(ScrambleParserNum, i)
}

func (s *SwapPositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwapPositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwapPositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterSwapPosition(s)
	}
}

func (s *SwapPositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitSwapPosition(s)
	}
}

func (p *ScrambleParser) SwapPosition() (localctx ISwapPositionContext) {
	localctx = NewSwapPositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ScrambleParserRULE_swapPosition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ScrambleParserT__0)
	}
	{
		p.SetState(39)
		p.Match(ScrambleParserNum)
	}
	{
		p.SetState(40)
		p.Match(ScrambleParserT__1)
	}
	{
		p.SetState(41)
		p.Match(ScrambleParserNum)
	}

	return localctx
}

// ISwapLetterContext is an interface to support dynamic dispatch.
type ISwapLetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwapLetterContext differentiates from other interfaces.
	IsSwapLetterContext()
}

type SwapLetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwapLetterContext() *SwapLetterContext {
	var p = new(SwapLetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_swapLetter
	return p
}

func (*SwapLetterContext) IsSwapLetterContext() {}

func NewSwapLetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwapLetterContext {
	var p = new(SwapLetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_swapLetter

	return p
}

func (s *SwapLetterContext) GetParser() antlr.Parser { return s.parser }

func (s *SwapLetterContext) AllLetter() []antlr.TerminalNode {
	return s.GetTokens(ScrambleParserLetter)
}

func (s *SwapLetterContext) Letter(i int) antlr.TerminalNode {
	return s.GetToken(ScrambleParserLetter, i)
}

func (s *SwapLetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwapLetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwapLetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterSwapLetter(s)
	}
}

func (s *SwapLetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitSwapLetter(s)
	}
}

func (p *ScrambleParser) SwapLetter() (localctx ISwapLetterContext) {
	localctx = NewSwapLetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ScrambleParserRULE_swapLetter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(ScrambleParserT__2)
	}
	{
		p.SetState(44)
		p.Match(ScrambleParserLetter)
	}
	{
		p.SetState(45)
		p.Match(ScrambleParserT__3)
	}
	{
		p.SetState(46)
		p.Match(ScrambleParserLetter)
	}

	return localctx
}

// IRotateLeftContext is an interface to support dynamic dispatch.
type IRotateLeftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRotateLeftContext differentiates from other interfaces.
	IsRotateLeftContext()
}

type RotateLeftContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRotateLeftContext() *RotateLeftContext {
	var p = new(RotateLeftContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_rotateLeft
	return p
}

func (*RotateLeftContext) IsRotateLeftContext() {}

func NewRotateLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RotateLeftContext {
	var p = new(RotateLeftContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_rotateLeft

	return p
}

func (s *RotateLeftContext) GetParser() antlr.Parser { return s.parser }

func (s *RotateLeftContext) Num() antlr.TerminalNode {
	return s.GetToken(ScrambleParserNum, 0)
}

func (s *RotateLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RotateLeftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RotateLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterRotateLeft(s)
	}
}

func (s *RotateLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitRotateLeft(s)
	}
}

func (p *ScrambleParser) RotateLeft() (localctx IRotateLeftContext) {
	localctx = NewRotateLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ScrambleParserRULE_rotateLeft)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(ScrambleParserT__4)
	}
	{
		p.SetState(49)
		p.Match(ScrambleParserNum)
	}
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ScrambleParserT__5 || _la == ScrambleParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRotateRightContext is an interface to support dynamic dispatch.
type IRotateRightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRotateRightContext differentiates from other interfaces.
	IsRotateRightContext()
}

type RotateRightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRotateRightContext() *RotateRightContext {
	var p = new(RotateRightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_rotateRight
	return p
}

func (*RotateRightContext) IsRotateRightContext() {}

func NewRotateRightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RotateRightContext {
	var p = new(RotateRightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_rotateRight

	return p
}

func (s *RotateRightContext) GetParser() antlr.Parser { return s.parser }

func (s *RotateRightContext) Num() antlr.TerminalNode {
	return s.GetToken(ScrambleParserNum, 0)
}

func (s *RotateRightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RotateRightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RotateRightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterRotateRight(s)
	}
}

func (s *RotateRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitRotateRight(s)
	}
}

func (p *ScrambleParser) RotateRight() (localctx IRotateRightContext) {
	localctx = NewRotateRightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ScrambleParserRULE_rotateRight)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(ScrambleParserT__7)
	}
	{
		p.SetState(53)
		p.Match(ScrambleParserNum)
	}
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ScrambleParserT__5 || _la == ScrambleParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRotateLetterContext is an interface to support dynamic dispatch.
type IRotateLetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRotateLetterContext differentiates from other interfaces.
	IsRotateLetterContext()
}

type RotateLetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRotateLetterContext() *RotateLetterContext {
	var p = new(RotateLetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_rotateLetter
	return p
}

func (*RotateLetterContext) IsRotateLetterContext() {}

func NewRotateLetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RotateLetterContext {
	var p = new(RotateLetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_rotateLetter

	return p
}

func (s *RotateLetterContext) GetParser() antlr.Parser { return s.parser }

func (s *RotateLetterContext) Letter() antlr.TerminalNode {
	return s.GetToken(ScrambleParserLetter, 0)
}

func (s *RotateLetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RotateLetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RotateLetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterRotateLetter(s)
	}
}

func (s *RotateLetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitRotateLetter(s)
	}
}

func (p *ScrambleParser) RotateLetter() (localctx IRotateLetterContext) {
	localctx = NewRotateLetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ScrambleParserRULE_rotateLetter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(ScrambleParserT__8)
	}
	{
		p.SetState(57)
		p.Match(ScrambleParserLetter)
	}

	return localctx
}

// IReverseContext is an interface to support dynamic dispatch.
type IReverseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReverseContext differentiates from other interfaces.
	IsReverseContext()
}

type ReverseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReverseContext() *ReverseContext {
	var p = new(ReverseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_reverse
	return p
}

func (*ReverseContext) IsReverseContext() {}

func NewReverseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReverseContext {
	var p = new(ReverseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_reverse

	return p
}

func (s *ReverseContext) GetParser() antlr.Parser { return s.parser }

func (s *ReverseContext) AllNum() []antlr.TerminalNode {
	return s.GetTokens(ScrambleParserNum)
}

func (s *ReverseContext) Num(i int) antlr.TerminalNode {
	return s.GetToken(ScrambleParserNum, i)
}

func (s *ReverseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReverseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReverseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterReverse(s)
	}
}

func (s *ReverseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitReverse(s)
	}
}

func (p *ScrambleParser) Reverse() (localctx IReverseContext) {
	localctx = NewReverseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ScrambleParserRULE_reverse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(ScrambleParserT__9)
	}
	{
		p.SetState(60)
		p.Match(ScrambleParserNum)
	}
	{
		p.SetState(61)
		p.Match(ScrambleParserT__10)
	}
	{
		p.SetState(62)
		p.Match(ScrambleParserNum)
	}

	return localctx
}

// IMoveContext is an interface to support dynamic dispatch.
type IMoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoveContext differentiates from other interfaces.
	IsMoveContext()
}

type MoveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveContext() *MoveContext {
	var p = new(MoveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScrambleParserRULE_move
	return p
}

func (*MoveContext) IsMoveContext() {}

func NewMoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveContext {
	var p = new(MoveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScrambleParserRULE_move

	return p
}

func (s *MoveContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveContext) AllNum() []antlr.TerminalNode {
	return s.GetTokens(ScrambleParserNum)
}

func (s *MoveContext) Num(i int) antlr.TerminalNode {
	return s.GetToken(ScrambleParserNum, i)
}

func (s *MoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.EnterMove(s)
	}
}

func (s *MoveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScrambleListener); ok {
		listenerT.ExitMove(s)
	}
}

func (p *ScrambleParser) Move() (localctx IMoveContext) {
	localctx = NewMoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ScrambleParserRULE_move)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(ScrambleParserT__11)
	}
	{
		p.SetState(65)
		p.Match(ScrambleParserNum)
	}
	{
		p.SetState(66)
		p.Match(ScrambleParserT__12)
	}
	{
		p.SetState(67)
		p.Match(ScrambleParserNum)
	}

	return localctx
}
