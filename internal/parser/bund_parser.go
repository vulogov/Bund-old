// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 3, 3, 3, 5, 3, 23, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 29, 10, 4, 12, 4, 14, 4, 32, 11, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 7, 5, 38, 10, 5, 12, 5, 14, 5, 41, 11, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 5, 6, 48, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 54,
	10, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3, 3, 2, 9, 10, 2, 56, 2,
	17, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 35, 3, 2, 2, 2,
	10, 47, 3, 2, 2, 2, 12, 49, 3, 2, 2, 2, 14, 16, 5, 4, 3, 2, 15, 14, 3,
	2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18,
	3, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 23, 5, 6, 4, 2, 21, 23, 5, 8, 5,
	2, 22, 20, 3, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 25, 7,
	3, 2, 2, 25, 26, 7, 12, 2, 2, 26, 30, 7, 4, 2, 2, 27, 29, 5, 10, 6, 2,
	28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3,
	2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 34, 7, 5, 2, 2, 34,
	7, 3, 2, 2, 2, 35, 39, 7, 6, 2, 2, 36, 38, 5, 10, 6, 2, 37, 36, 3, 2, 2,
	2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42,
	3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 7, 7, 2, 2, 43, 9, 3, 2, 2, 2,
	44, 48, 5, 6, 4, 2, 45, 48, 5, 8, 5, 2, 46, 48, 5, 12, 7, 2, 47, 44, 3,
	2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 11, 3, 2, 2, 2, 49,
	53, 9, 2, 2, 2, 50, 51, 7, 8, 2, 2, 51, 52, 7, 12, 2, 2, 52, 54, 7, 7,
	2, 2, 53, 50, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 13, 3, 2, 2, 2, 8, 17,
	22, 30, 39, 47, 53,
}
var literalNames = []string{
	"", "'['", "':'", "';;'", "'('", "')'", "'.('", "", "", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "TRUE", "FALSE", "SLASH", "NAME", "COMMENT",
	"BLOCK_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"expressions", "root_term", "ns", "block", "term", "boolean_term",
}

type BundParser struct {
	*antlr.BaseParser
}

// NewBundParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BundParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundParser(input antlr.TokenStream) *BundParser {
	this := new(BundParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Bund.g4"

	return this
}

// BundParser tokens.
const (
	BundParserEOF           = antlr.TokenEOF
	BundParserT__0          = 1
	BundParserT__1          = 2
	BundParserT__2          = 3
	BundParserT__3          = 4
	BundParserT__4          = 5
	BundParserT__5          = 6
	BundParserTRUE          = 7
	BundParserFALSE         = 8
	BundParserSLASH         = 9
	BundParserNAME          = 10
	BundParserCOMMENT       = 11
	BundParserBLOCK_COMMENT = 12
	BundParserWS            = 13
	BundParserSHEBANG       = 14
)

// BundParser rules.
const (
	BundParserRULE_expressions  = 0
	BundParserRULE_root_term    = 1
	BundParserRULE_ns           = 2
	BundParserRULE_block        = 3
	BundParserRULE_term         = 4
	BundParserRULE_boolean_term = 5
)

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllRoot_term() []IRoot_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRoot_termContext)(nil)).Elem())
	var tst = make([]IRoot_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRoot_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Root_term(i int) IRoot_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoot_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRoot_termContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *BundParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BundParserRULE_expressions)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserT__0 || _la == BundParserT__3 {
		{
			p.SetState(12)
			p.Root_term()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRoot_termContext is an interface to support dynamic dispatch.
type IRoot_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoot_termContext differentiates from other interfaces.
	IsRoot_termContext()
}

type Root_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoot_termContext() *Root_termContext {
	var p = new(Root_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_root_term
	return p
}

func (*Root_termContext) IsRoot_termContext() {}

func NewRoot_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Root_termContext {
	var p = new(Root_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_root_term

	return p
}

func (s *Root_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Root_termContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Root_termContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Root_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Root_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Root_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterRoot_term(s)
	}
}

func (s *Root_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitRoot_term(s)
	}
}

func (p *BundParser) Root_term() (localctx IRoot_termContext) {
	localctx = NewRoot_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BundParserRULE_root_term)

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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(18)
			p.Ns()
		}

	case BundParserT__3:
		{
			p.SetState(19)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INsContext is an interface to support dynamic dispatch.
type INsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsNsContext differentiates from other interfaces.
	IsNsContext()
}

type NsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyNsContext() *NsContext {
	var p = new(NsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ns
	return p
}

func (*NsContext) IsNsContext() {}

func NewNsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NsContext {
	var p = new(NsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ns

	return p
}

func (s *NsContext) GetParser() antlr.Parser { return s.parser }

func (s *NsContext) GetName() antlr.Token { return s.name }

func (s *NsContext) SetName(v antlr.Token) { s.name = v }

func (s *NsContext) Get_term() ITermContext { return s._term }

func (s *NsContext) Set_term(v ITermContext) { s._term = v }

func (s *NsContext) GetBody() []ITermContext { return s.body }

func (s *NsContext) SetBody(v []ITermContext) { s.body = v }

func (s *NsContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *NsContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *NsContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *NsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterNs(s)
	}
}

func (s *NsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitNs(s)
	}
}

func (p *BundParser) Ns() (localctx INsContext) {
	localctx = NewNsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BundParserRULE_ns)
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
		p.SetState(22)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(23)

		var _m = p.Match(BundParserNAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(24)
		p.Match(BundParserT__1)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__3)|(1<<BundParserTRUE)|(1<<BundParserFALSE))) != 0 {
		{
			p.SetState(25)

			var _x = p.Term()

			localctx.(*NsContext)._term = _x
		}
		localctx.(*NsContext).body = append(localctx.(*NsContext).body, localctx.(*NsContext)._term)

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(BundParserT__2)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_term() ITermContext { return s._term }

func (s *BlockContext) Set_term(v ITermContext) { s._term = v }

func (s *BlockContext) GetBody() []ITermContext { return s.body }

func (s *BlockContext) SetBody(v []ITermContext) { s.body = v }

func (s *BlockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *BlockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BundParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_block)
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
		p.SetState(33)
		p.Match(BundParserT__3)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__3)|(1<<BundParserTRUE)|(1<<BundParserFALSE))) != 0 {
		{
			p.SetState(34)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).body = append(localctx.(*BlockContext).body, localctx.(*BlockContext)._term)

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(BundParserT__4)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *TermContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TermContext) Boolean_term() IBoolean_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_termContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BundParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BundParserRULE_term)

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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(42)
			p.Ns()
		}

	case BundParserT__3:
		{
			p.SetState(43)
			p.Block()
		}

	case BundParserTRUE, BundParserFALSE:
		{
			p.SetState(44)
			p.Boolean_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolean_termContext is an interface to support dynamic dispatch.
type IBoolean_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// GetFUNCTOR returns the FUNCTOR token.
	GetFUNCTOR() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// SetFUNCTOR sets the FUNCTOR token.
	SetFUNCTOR(antlr.Token)

	// IsBoolean_termContext differentiates from other interfaces.
	IsBoolean_termContext()
}

type Boolean_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyBoolean_termContext() *Boolean_termContext {
	var p = new(Boolean_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_boolean_term
	return p
}

func (*Boolean_termContext) IsBoolean_termContext() {}

func NewBoolean_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_termContext {
	var p = new(Boolean_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_boolean_term

	return p
}

func (s *Boolean_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Boolean_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Boolean_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Boolean_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Boolean_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BundParserTRUE, 0)
}

func (s *Boolean_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BundParserFALSE, 0)
}

func (s *Boolean_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Boolean_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBoolean_term(s)
	}
}

func (s *Boolean_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBoolean_term(s)
	}
}

func (p *BundParser) Boolean_term() (localctx IBoolean_termContext) {
	localctx = NewBoolean_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_boolean_term)
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
		p.SetState(47)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Boolean_termContext).VALUE = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserTRUE || _la == BundParserFALSE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Boolean_termContext).VALUE = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__5 {
		{
			p.SetState(48)
			p.Match(BundParserT__5)
		}
		{
			p.SetState(49)

			var _m = p.Match(BundParserNAME)

			localctx.(*Boolean_termContext).FUNCTOR = _m
		}
		{
			p.SetState(50)
			p.Match(BundParserT__4)
		}

	}

	return localctx
}
