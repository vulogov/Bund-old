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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 169,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 7, 2, 42, 10, 2, 12, 2, 14, 2, 45,
	11, 2, 3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 55, 10,
	4, 12, 4, 14, 4, 58, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 64, 10, 5, 12,
	5, 14, 5, 67, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 86, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 102, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 108, 10,
	9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 120, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 126, 10, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 132, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	5, 15, 138, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 144, 10, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 150, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 7, 20, 158, 10, 20, 12, 20, 14, 20, 161, 11, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 167, 10, 20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 6, 4, 2, 19,
	19, 21, 21, 4, 2, 17, 17, 21, 21, 3, 2, 15, 16, 3, 2, 22, 23, 2, 184, 2,
	43, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 61, 3, 2, 2, 2,
	10, 85, 3, 2, 2, 2, 12, 95, 3, 2, 2, 2, 14, 97, 3, 2, 2, 2, 16, 103, 3,
	2, 2, 2, 18, 109, 3, 2, 2, 2, 20, 112, 3, 2, 2, 2, 22, 115, 3, 2, 2, 2,
	24, 121, 3, 2, 2, 2, 26, 127, 3, 2, 2, 2, 28, 133, 3, 2, 2, 2, 30, 139,
	3, 2, 2, 2, 32, 145, 3, 2, 2, 2, 34, 151, 3, 2, 2, 2, 36, 153, 3, 2, 2,
	2, 38, 155, 3, 2, 2, 2, 40, 42, 5, 4, 3, 2, 41, 40, 3, 2, 2, 2, 42, 45,
	3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 3, 3, 2, 2, 2,
	45, 43, 3, 2, 2, 2, 46, 49, 5, 6, 4, 2, 47, 49, 5, 8, 5, 2, 48, 46, 3,
	2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 5, 3, 2, 2, 2, 50, 51, 7, 3, 2, 2, 51,
	52, 7, 21, 2, 2, 52, 56, 7, 22, 2, 2, 53, 55, 5, 10, 6, 2, 54, 53, 3, 2,
	2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59,
	3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7, 4, 2, 2, 60, 7, 3, 2, 2, 2,
	61, 65, 7, 5, 2, 2, 62, 64, 5, 10, 6, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3,
	2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67,
	65, 3, 2, 2, 2, 68, 69, 7, 6, 2, 2, 69, 9, 3, 2, 2, 2, 70, 86, 5, 6, 4,
	2, 71, 86, 5, 8, 5, 2, 72, 86, 5, 14, 8, 2, 73, 86, 5, 16, 9, 2, 74, 86,
	5, 18, 10, 2, 75, 86, 5, 20, 11, 2, 76, 86, 5, 22, 12, 2, 77, 86, 5, 24,
	13, 2, 78, 86, 5, 26, 14, 2, 79, 86, 5, 28, 15, 2, 80, 86, 5, 30, 16, 2,
	81, 86, 5, 32, 17, 2, 82, 86, 5, 38, 20, 2, 83, 86, 5, 34, 18, 2, 84, 86,
	5, 36, 19, 2, 85, 70, 3, 2, 2, 2, 85, 71, 3, 2, 2, 2, 85, 72, 3, 2, 2,
	2, 85, 73, 3, 2, 2, 2, 85, 74, 3, 2, 2, 2, 85, 75, 3, 2, 2, 2, 85, 76,
	3, 2, 2, 2, 85, 77, 3, 2, 2, 2, 85, 78, 3, 2, 2, 2, 85, 79, 3, 2, 2, 2,
	85, 80, 3, 2, 2, 2, 85, 81, 3, 2, 2, 2, 85, 82, 3, 2, 2, 2, 85, 83, 3,
	2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 96, 5, 22, 12, 2,
	88, 96, 5, 24, 13, 2, 89, 96, 5, 26, 14, 2, 90, 96, 5, 28, 15, 2, 91, 96,
	5, 30, 16, 2, 92, 96, 5, 14, 8, 2, 93, 96, 5, 16, 9, 2, 94, 96, 5, 36,
	19, 2, 95, 87, 3, 2, 2, 2, 95, 88, 3, 2, 2, 2, 95, 89, 3, 2, 2, 2, 95,
	90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2,
	2, 95, 94, 3, 2, 2, 2, 96, 13, 3, 2, 2, 2, 97, 101, 9, 2, 2, 2, 98, 99,
	7, 7, 2, 2, 99, 100, 9, 3, 2, 2, 100, 102, 7, 6, 2, 2, 101, 98, 3, 2, 2,
	2, 101, 102, 3, 2, 2, 2, 102, 15, 3, 2, 2, 2, 103, 107, 7, 18, 2, 2, 104,
	105, 7, 7, 2, 2, 105, 106, 9, 3, 2, 2, 106, 108, 7, 6, 2, 2, 107, 104,
	3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 17, 3, 2, 2, 2, 109, 110, 7, 8,
	2, 2, 110, 111, 9, 2, 2, 2, 111, 19, 3, 2, 2, 2, 112, 113, 7, 8, 2, 2,
	113, 114, 7, 18, 2, 2, 114, 21, 3, 2, 2, 2, 115, 119, 9, 4, 2, 2, 116,
	117, 7, 7, 2, 2, 117, 118, 9, 2, 2, 2, 118, 120, 7, 6, 2, 2, 119, 116,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 23, 3, 2, 2, 2, 121, 125, 7, 10,
	2, 2, 122, 123, 7, 7, 2, 2, 123, 124, 9, 2, 2, 2, 124, 126, 7, 6, 2, 2,
	125, 122, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 25, 3, 2, 2, 2, 127, 131,
	7, 12, 2, 2, 128, 129, 7, 7, 2, 2, 129, 130, 9, 2, 2, 2, 130, 132, 7, 6,
	2, 2, 131, 128, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 27, 3, 2, 2, 2,
	133, 137, 7, 13, 2, 2, 134, 135, 7, 7, 2, 2, 135, 136, 9, 2, 2, 2, 136,
	138, 7, 6, 2, 2, 137, 134, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 29, 3,
	2, 2, 2, 139, 143, 7, 14, 2, 2, 140, 141, 7, 7, 2, 2, 141, 142, 9, 2, 2,
	2, 142, 144, 7, 6, 2, 2, 143, 140, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144,
	31, 3, 2, 2, 2, 145, 149, 7, 24, 2, 2, 146, 147, 7, 7, 2, 2, 147, 148,
	9, 2, 2, 2, 148, 150, 7, 6, 2, 2, 149, 146, 3, 2, 2, 2, 149, 150, 3, 2,
	2, 2, 150, 33, 3, 2, 2, 2, 151, 152, 9, 5, 2, 2, 152, 35, 3, 2, 2, 2, 153,
	154, 7, 25, 2, 2, 154, 37, 3, 2, 2, 2, 155, 159, 7, 9, 2, 2, 156, 158,
	5, 12, 7, 2, 157, 156, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2,
	2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2,
	162, 166, 7, 6, 2, 2, 163, 164, 7, 7, 2, 2, 164, 165, 9, 3, 2, 2, 165,
	167, 7, 6, 2, 2, 166, 163, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 39, 3,
	2, 2, 2, 18, 43, 48, 56, 65, 85, 95, 101, 107, 119, 125, 131, 137, 143,
	149, 159, 166,
}
var literalNames = []string{
	"", "'['", "';;'", "'('", "')'", "'.('", "'`'", "'(*'", "", "", "", "",
	"", "", "", "", "", "", "'/'", "", "':'", "';'", "", "'|'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER",
	"STRING", "COMPLEX_NUMBER", "TRUE", "FALSE", "SYS", "CMD", "SYSF", "SLASH",
	"NAME", "TOBEGIN", "TOEND", "GLOB", "SEPARATE", "COMMENT", "BLOCK_COMMENT",
	"WS", "SHEBANG",
}

var ruleNames = []string{
	"expressions", "root_term", "ns", "block", "term", "data", "call_term",
	"operator_term", "ref_call_term", "ref_operator_term", "boolean_term",
	"integer_term", "float_term", "string_term", "complex_term", "glob_term",
	"mode_term", "separate_term", "datablock_term",
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
	BundParserEOF             = antlr.TokenEOF
	BundParserT__0            = 1
	BundParserT__1            = 2
	BundParserT__2            = 3
	BundParserT__3            = 4
	BundParserT__4            = 5
	BundParserT__5            = 6
	BundParserT__6            = 7
	BundParserINTEGER         = 8
	BundParserDECIMAL_INTEGER = 9
	BundParserFLOAT_NUMBER    = 10
	BundParserSTRING          = 11
	BundParserCOMPLEX_NUMBER  = 12
	BundParserTRUE            = 13
	BundParserFALSE           = 14
	BundParserSYS             = 15
	BundParserCMD             = 16
	BundParserSYSF            = 17
	BundParserSLASH           = 18
	BundParserNAME            = 19
	BundParserTOBEGIN         = 20
	BundParserTOEND           = 21
	BundParserGLOB            = 22
	BundParserSEPARATE        = 23
	BundParserCOMMENT         = 24
	BundParserBLOCK_COMMENT   = 25
	BundParserWS              = 26
	BundParserSHEBANG         = 27
)

// BundParser rules.
const (
	BundParserRULE_expressions       = 0
	BundParserRULE_root_term         = 1
	BundParserRULE_ns                = 2
	BundParserRULE_block             = 3
	BundParserRULE_term              = 4
	BundParserRULE_data              = 5
	BundParserRULE_call_term         = 6
	BundParserRULE_operator_term     = 7
	BundParserRULE_ref_call_term     = 8
	BundParserRULE_ref_operator_term = 9
	BundParserRULE_boolean_term      = 10
	BundParserRULE_integer_term      = 11
	BundParserRULE_float_term        = 12
	BundParserRULE_string_term       = 13
	BundParserRULE_complex_term      = 14
	BundParserRULE_glob_term         = 15
	BundParserRULE_mode_term         = 16
	BundParserRULE_separate_term     = 17
	BundParserRULE_datablock_term    = 18
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserT__0 || _la == BundParserT__2 {
		{
			p.SetState(38)
			p.Root_term()
		}

		p.SetState(43)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(44)
			p.Ns()
		}

	case BundParserT__2:
		{
			p.SetState(45)
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

func (s *NsContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

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
		p.SetState(48)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(49)

		var _m = p.Match(BundParserNAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(50)
		p.Match(BundParserTOBEGIN)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(51)

			var _x = p.Term()

			localctx.(*NsContext)._term = _x
		}
		localctx.(*NsContext).body = append(localctx.(*NsContext).body, localctx.(*NsContext)._term)

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(BundParserT__1)
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
		p.SetState(59)
		p.Match(BundParserT__2)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(60)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).body = append(localctx.(*BlockContext).body, localctx.(*BlockContext)._term)

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(BundParserT__3)
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

func (s *TermContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *TermContext) Operator_term() IOperator_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperator_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperator_termContext)
}

func (s *TermContext) Ref_call_term() IRef_call_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRef_call_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRef_call_termContext)
}

func (s *TermContext) Ref_operator_term() IRef_operator_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRef_operator_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRef_operator_termContext)
}

func (s *TermContext) Boolean_term() IBoolean_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_termContext)
}

func (s *TermContext) Integer_term() IInteger_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_termContext)
}

func (s *TermContext) Float_term() IFloat_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_termContext)
}

func (s *TermContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *TermContext) Complex_term() IComplex_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplex_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplex_termContext)
}

func (s *TermContext) Glob_term() IGlob_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlob_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlob_termContext)
}

func (s *TermContext) Datablock_term() IDatablock_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatablock_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatablock_termContext)
}

func (s *TermContext) Mode_term() IMode_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMode_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMode_termContext)
}

func (s *TermContext) Separate_term() ISeparate_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparate_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparate_termContext)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(68)
			p.Ns()
		}

	case 2:
		{
			p.SetState(69)
			p.Block()
		}

	case 3:
		{
			p.SetState(70)
			p.Call_term()
		}

	case 4:
		{
			p.SetState(71)
			p.Operator_term()
		}

	case 5:
		{
			p.SetState(72)
			p.Ref_call_term()
		}

	case 6:
		{
			p.SetState(73)
			p.Ref_operator_term()
		}

	case 7:
		{
			p.SetState(74)
			p.Boolean_term()
		}

	case 8:
		{
			p.SetState(75)
			p.Integer_term()
		}

	case 9:
		{
			p.SetState(76)
			p.Float_term()
		}

	case 10:
		{
			p.SetState(77)
			p.String_term()
		}

	case 11:
		{
			p.SetState(78)
			p.Complex_term()
		}

	case 12:
		{
			p.SetState(79)
			p.Glob_term()
		}

	case 13:
		{
			p.SetState(80)
			p.Datablock_term()
		}

	case 14:
		{
			p.SetState(81)
			p.Mode_term()
		}

	case 15:
		{
			p.SetState(82)
			p.Separate_term()
		}

	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) Boolean_term() IBoolean_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_termContext)
}

func (s *DataContext) Integer_term() IInteger_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_termContext)
}

func (s *DataContext) Float_term() IFloat_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_termContext)
}

func (s *DataContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *DataContext) Complex_term() IComplex_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplex_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplex_termContext)
}

func (s *DataContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *DataContext) Operator_term() IOperator_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperator_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperator_termContext)
}

func (s *DataContext) Separate_term() ISeparate_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparate_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparate_termContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *BundParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_data)

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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserTRUE, BundParserFALSE:
		{
			p.SetState(85)
			p.Boolean_term()
		}

	case BundParserINTEGER:
		{
			p.SetState(86)
			p.Integer_term()
		}

	case BundParserFLOAT_NUMBER:
		{
			p.SetState(87)
			p.Float_term()
		}

	case BundParserSTRING:
		{
			p.SetState(88)
			p.String_term()
		}

	case BundParserCOMPLEX_NUMBER:
		{
			p.SetState(89)
			p.Complex_term()
		}

	case BundParserSYSF, BundParserNAME:
		{
			p.SetState(90)
			p.Call_term()
		}

	case BundParserCMD:
		{
			p.SetState(91)
			p.Operator_term()
		}

	case BundParserSEPARATE:
		{
			p.SetState(92)
			p.Separate_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICall_termContext is an interface to support dynamic dispatch.
type ICall_termContext interface {
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

	// IsCall_termContext differentiates from other interfaces.
	IsCall_termContext()
}

type Call_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyCall_termContext() *Call_termContext {
	var p = new(Call_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_call_term
	return p
}

func (*Call_termContext) IsCall_termContext() {}

func NewCall_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_termContext {
	var p = new(Call_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_call_term

	return p
}

func (s *Call_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Call_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Call_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Call_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Call_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *Call_termContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(BundParserNAME)
}

func (s *Call_termContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(BundParserNAME, i)
}

func (s *Call_termContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Call_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCall_term(s)
	}
}

func (s *Call_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCall_term(s)
	}
}

func (p *BundParser) Call_term() (localctx ICall_termContext) {
	localctx = NewCall_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_call_term)
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
		p.SetState(95)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Call_termContext).VALUE = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserSYSF || _la == BundParserNAME) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Call_termContext).VALUE = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(96)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(97)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Call_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYS || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Call_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(98)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IOperator_termContext is an interface to support dynamic dispatch.
type IOperator_termContext interface {
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

	// IsOperator_termContext differentiates from other interfaces.
	IsOperator_termContext()
}

type Operator_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyOperator_termContext() *Operator_termContext {
	var p = new(Operator_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_operator_term
	return p
}

func (*Operator_termContext) IsOperator_termContext() {}

func NewOperator_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_termContext {
	var p = new(Operator_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_operator_term

	return p
}

func (s *Operator_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Operator_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Operator_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Operator_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Operator_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Operator_termContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Operator_termContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Operator_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Operator_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterOperator_term(s)
	}
}

func (s *Operator_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitOperator_term(s)
	}
}

func (p *BundParser) Operator_term() (localctx IOperator_termContext) {
	localctx = NewOperator_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BundParserRULE_operator_term)
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
		p.SetState(101)

		var _m = p.Match(BundParserCMD)

		localctx.(*Operator_termContext).VALUE = _m
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(102)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(103)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Operator_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYS || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Operator_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(104)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IRef_call_termContext is an interface to support dynamic dispatch.
type IRef_call_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsRef_call_termContext differentiates from other interfaces.
	IsRef_call_termContext()
}

type Ref_call_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyRef_call_termContext() *Ref_call_termContext {
	var p = new(Ref_call_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ref_call_term
	return p
}

func (*Ref_call_termContext) IsRef_call_termContext() {}

func NewRef_call_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ref_call_termContext {
	var p = new(Ref_call_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ref_call_term

	return p
}

func (s *Ref_call_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ref_call_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Ref_call_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Ref_call_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *Ref_call_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Ref_call_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ref_call_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ref_call_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterRef_call_term(s)
	}
}

func (s *Ref_call_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitRef_call_term(s)
	}
}

func (p *BundParser) Ref_call_term() (localctx IRef_call_termContext) {
	localctx = NewRef_call_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BundParserRULE_ref_call_term)
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
		p.SetState(107)
		p.Match(BundParserT__5)
	}
	{
		p.SetState(108)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Ref_call_termContext).VALUE = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserSYSF || _la == BundParserNAME) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Ref_call_termContext).VALUE = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRef_operator_termContext is an interface to support dynamic dispatch.
type IRef_operator_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsRef_operator_termContext differentiates from other interfaces.
	IsRef_operator_termContext()
}

type Ref_operator_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyRef_operator_termContext() *Ref_operator_termContext {
	var p = new(Ref_operator_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ref_operator_term
	return p
}

func (*Ref_operator_termContext) IsRef_operator_termContext() {}

func NewRef_operator_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ref_operator_termContext {
	var p = new(Ref_operator_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ref_operator_term

	return p
}

func (s *Ref_operator_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ref_operator_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Ref_operator_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Ref_operator_termContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Ref_operator_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ref_operator_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ref_operator_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterRef_operator_term(s)
	}
}

func (s *Ref_operator_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitRef_operator_term(s)
	}
}

func (p *BundParser) Ref_operator_term() (localctx IRef_operator_termContext) {
	localctx = NewRef_operator_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BundParserRULE_ref_operator_term)

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
		p.SetState(110)
		p.Match(BundParserT__5)
	}
	{
		p.SetState(111)

		var _m = p.Match(BundParserCMD)

		localctx.(*Ref_operator_termContext).VALUE = _m
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

func (s *Boolean_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
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
	p.EnterRule(localctx, 20, BundParserRULE_boolean_term)
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
		p.SetState(113)

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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(114)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(115)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Boolean_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYSF || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Boolean_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(116)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IInteger_termContext is an interface to support dynamic dispatch.
type IInteger_termContext interface {
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

	// IsInteger_termContext differentiates from other interfaces.
	IsInteger_termContext()
}

type Integer_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyInteger_termContext() *Integer_termContext {
	var p = new(Integer_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_integer_term
	return p
}

func (*Integer_termContext) IsInteger_termContext() {}

func NewInteger_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_termContext {
	var p = new(Integer_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_integer_term

	return p
}

func (s *Integer_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Integer_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Integer_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Integer_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Integer_termContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserINTEGER, 0)
}

func (s *Integer_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *Integer_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Integer_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterInteger_term(s)
	}
}

func (s *Integer_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitInteger_term(s)
	}
}

func (p *BundParser) Integer_term() (localctx IInteger_termContext) {
	localctx = NewInteger_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BundParserRULE_integer_term)
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
		p.SetState(119)

		var _m = p.Match(BundParserINTEGER)

		localctx.(*Integer_termContext).VALUE = _m
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(120)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(121)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Integer_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYSF || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Integer_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(122)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IFloat_termContext is an interface to support dynamic dispatch.
type IFloat_termContext interface {
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

	// IsFloat_termContext differentiates from other interfaces.
	IsFloat_termContext()
}

type Float_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyFloat_termContext() *Float_termContext {
	var p = new(Float_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_float_term
	return p
}

func (*Float_termContext) IsFloat_termContext() {}

func NewFloat_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_termContext {
	var p = new(Float_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_float_term

	return p
}

func (s *Float_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Float_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Float_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Float_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Float_termContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserFLOAT_NUMBER, 0)
}

func (s *Float_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *Float_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Float_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFloat_term(s)
	}
}

func (s *Float_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFloat_term(s)
	}
}

func (p *BundParser) Float_term() (localctx IFloat_termContext) {
	localctx = NewFloat_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BundParserRULE_float_term)
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
		p.SetState(125)

		var _m = p.Match(BundParserFLOAT_NUMBER)

		localctx.(*Float_termContext).VALUE = _m
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(126)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(127)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Float_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYSF || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Float_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(128)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IString_termContext is an interface to support dynamic dispatch.
type IString_termContext interface {
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

	// IsString_termContext differentiates from other interfaces.
	IsString_termContext()
}

type String_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyString_termContext() *String_termContext {
	var p = new(String_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_string_term
	return p
}

func (*String_termContext) IsString_termContext() {}

func NewString_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_termContext {
	var p = new(String_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_string_term

	return p
}

func (s *String_termContext) GetParser() antlr.Parser { return s.parser }

func (s *String_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *String_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *String_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *String_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *String_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(BundParserSTRING, 0)
}

func (s *String_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *String_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *String_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterString_term(s)
	}
}

func (s *String_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitString_term(s)
	}
}

func (p *BundParser) String_term() (localctx IString_termContext) {
	localctx = NewString_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BundParserRULE_string_term)
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
		p.SetState(131)

		var _m = p.Match(BundParserSTRING)

		localctx.(*String_termContext).VALUE = _m
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(132)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(133)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*String_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYSF || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*String_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(134)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IComplex_termContext is an interface to support dynamic dispatch.
type IComplex_termContext interface {
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

	// IsComplex_termContext differentiates from other interfaces.
	IsComplex_termContext()
}

type Complex_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyComplex_termContext() *Complex_termContext {
	var p = new(Complex_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_complex_term
	return p
}

func (*Complex_termContext) IsComplex_termContext() {}

func NewComplex_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Complex_termContext {
	var p = new(Complex_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_complex_term

	return p
}

func (s *Complex_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Complex_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Complex_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Complex_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Complex_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Complex_termContext) COMPLEX_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserCOMPLEX_NUMBER, 0)
}

func (s *Complex_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *Complex_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Complex_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Complex_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Complex_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterComplex_term(s)
	}
}

func (s *Complex_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitComplex_term(s)
	}
}

func (p *BundParser) Complex_term() (localctx IComplex_termContext) {
	localctx = NewComplex_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BundParserRULE_complex_term)
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
		p.SetState(137)

		var _m = p.Match(BundParserCOMPLEX_NUMBER)

		localctx.(*Complex_termContext).VALUE = _m
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(138)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(139)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Complex_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYSF || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Complex_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(140)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IGlob_termContext is an interface to support dynamic dispatch.
type IGlob_termContext interface {
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

	// IsGlob_termContext differentiates from other interfaces.
	IsGlob_termContext()
}

type Glob_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	VALUE   antlr.Token
	FUNCTOR antlr.Token
}

func NewEmptyGlob_termContext() *Glob_termContext {
	var p = new(Glob_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_glob_term
	return p
}

func (*Glob_termContext) IsGlob_termContext() {}

func NewGlob_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Glob_termContext {
	var p = new(Glob_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_glob_term

	return p
}

func (s *Glob_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Glob_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Glob_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Glob_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Glob_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Glob_termContext) GLOB() antlr.TerminalNode {
	return s.GetToken(BundParserGLOB, 0)
}

func (s *Glob_termContext) SYSF() antlr.TerminalNode {
	return s.GetToken(BundParserSYSF, 0)
}

func (s *Glob_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Glob_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Glob_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Glob_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterGlob_term(s)
	}
}

func (s *Glob_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitGlob_term(s)
	}
}

func (p *BundParser) Glob_term() (localctx IGlob_termContext) {
	localctx = NewGlob_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BundParserRULE_glob_term)
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
		p.SetState(143)

		var _m = p.Match(BundParserGLOB)

		localctx.(*Glob_termContext).VALUE = _m
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(144)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(145)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Glob_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYSF || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Glob_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IMode_termContext is an interface to support dynamic dispatch.
type IMode_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsMode_termContext differentiates from other interfaces.
	IsMode_termContext()
}

type Mode_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyMode_termContext() *Mode_termContext {
	var p = new(Mode_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_mode_term
	return p
}

func (*Mode_termContext) IsMode_termContext() {}

func NewMode_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mode_termContext {
	var p = new(Mode_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_mode_term

	return p
}

func (s *Mode_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Mode_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Mode_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Mode_termContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

func (s *Mode_termContext) TOEND() antlr.TerminalNode {
	return s.GetToken(BundParserTOEND, 0)
}

func (s *Mode_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mode_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mode_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterMode_term(s)
	}
}

func (s *Mode_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitMode_term(s)
	}
}

func (p *BundParser) Mode_term() (localctx IMode_termContext) {
	localctx = NewMode_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BundParserRULE_mode_term)
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
		p.SetState(149)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Mode_termContext).VALUE = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserTOBEGIN || _la == BundParserTOEND) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Mode_termContext).VALUE = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISeparate_termContext is an interface to support dynamic dispatch.
type ISeparate_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsSeparate_termContext differentiates from other interfaces.
	IsSeparate_termContext()
}

type Separate_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptySeparate_termContext() *Separate_termContext {
	var p = new(Separate_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_separate_term
	return p
}

func (*Separate_termContext) IsSeparate_termContext() {}

func NewSeparate_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Separate_termContext {
	var p = new(Separate_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_separate_term

	return p
}

func (s *Separate_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Separate_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Separate_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Separate_termContext) SEPARATE() antlr.TerminalNode {
	return s.GetToken(BundParserSEPARATE, 0)
}

func (s *Separate_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Separate_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Separate_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterSeparate_term(s)
	}
}

func (s *Separate_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitSeparate_term(s)
	}
}

func (p *BundParser) Separate_term() (localctx ISeparate_termContext) {
	localctx = NewSeparate_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BundParserRULE_separate_term)

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
		p.SetState(151)

		var _m = p.Match(BundParserSEPARATE)

		localctx.(*Separate_termContext).VALUE = _m
	}

	return localctx
}

// IDatablock_termContext is an interface to support dynamic dispatch.
type IDatablock_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFUNCTOR returns the FUNCTOR token.
	GetFUNCTOR() antlr.Token

	// SetFUNCTOR sets the FUNCTOR token.
	SetFUNCTOR(antlr.Token)

	// Get_data returns the _data rule contexts.
	Get_data() IDataContext

	// Set_data sets the _data rule contexts.
	Set_data(IDataContext)

	// GetBody returns the body rule context list.
	GetBody() []IDataContext

	// SetBody sets the body rule context list.
	SetBody([]IDataContext)

	// IsDatablock_termContext differentiates from other interfaces.
	IsDatablock_termContext()
}

type Datablock_termContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_data   IDataContext
	body    []IDataContext
	FUNCTOR antlr.Token
}

func NewEmptyDatablock_termContext() *Datablock_termContext {
	var p = new(Datablock_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_datablock_term
	return p
}

func (*Datablock_termContext) IsDatablock_termContext() {}

func NewDatablock_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datablock_termContext {
	var p = new(Datablock_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_datablock_term

	return p
}

func (s *Datablock_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Datablock_termContext) GetFUNCTOR() antlr.Token { return s.FUNCTOR }

func (s *Datablock_termContext) SetFUNCTOR(v antlr.Token) { s.FUNCTOR = v }

func (s *Datablock_termContext) Get_data() IDataContext { return s._data }

func (s *Datablock_termContext) Set_data(v IDataContext) { s._data = v }

func (s *Datablock_termContext) GetBody() []IDataContext { return s.body }

func (s *Datablock_termContext) SetBody(v []IDataContext) { s.body = v }

func (s *Datablock_termContext) AllData() []IDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataContext)(nil)).Elem())
	var tst = make([]IDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataContext)
		}
	}

	return tst
}

func (s *Datablock_termContext) Data(i int) IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *Datablock_termContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Datablock_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Datablock_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datablock_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datablock_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDatablock_term(s)
	}
}

func (s *Datablock_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDatablock_term(s)
	}
}

func (p *BundParser) Datablock_term() (localctx IDatablock_termContext) {
	localctx = NewDatablock_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BundParserRULE_datablock_term)
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
		p.SetState(153)
		p.Match(BundParserT__6)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(154)

			var _x = p.Data()

			localctx.(*Datablock_termContext)._data = _x
		}
		localctx.(*Datablock_termContext).body = append(localctx.(*Datablock_termContext).body, localctx.(*Datablock_termContext)._data)

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
		p.Match(BundParserT__3)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(161)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(162)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Datablock_termContext).FUNCTOR = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == BundParserSYS || _la == BundParserNAME) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Datablock_termContext).FUNCTOR = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(163)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}
