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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 193,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 7, 2,
	46, 10, 2, 12, 2, 14, 2, 49, 11, 2, 3, 3, 3, 3, 5, 3, 53, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 59, 10, 4, 12, 4, 14, 4, 62, 11, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 7, 5, 68, 10, 5, 12, 5, 14, 5, 71, 11, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 92, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 103, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	109, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 115, 10, 9, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 127, 10, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 133, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 5, 14, 139, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 145, 10, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 151, 10, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 157, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 7, 20,
	165, 10, 20, 12, 20, 14, 20, 168, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5,
	20, 174, 10, 20, 3, 21, 3, 21, 6, 21, 178, 10, 21, 13, 21, 14, 21, 179,
	3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 186, 10, 22, 12, 22, 14, 22, 189, 11,
	22, 3, 22, 3, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 7, 4, 2, 22, 22, 24, 24,
	4, 2, 20, 20, 24, 24, 3, 2, 18, 19, 3, 2, 25, 26, 3, 2, 11, 12, 2, 211,
	2, 47, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 6, 54, 3, 2, 2, 2, 8, 65, 3, 2, 2,
	2, 10, 91, 3, 2, 2, 2, 12, 102, 3, 2, 2, 2, 14, 104, 3, 2, 2, 2, 16, 110,
	3, 2, 2, 2, 18, 116, 3, 2, 2, 2, 20, 119, 3, 2, 2, 2, 22, 122, 3, 2, 2,
	2, 24, 128, 3, 2, 2, 2, 26, 134, 3, 2, 2, 2, 28, 140, 3, 2, 2, 2, 30, 146,
	3, 2, 2, 2, 32, 152, 3, 2, 2, 2, 34, 158, 3, 2, 2, 2, 36, 160, 3, 2, 2,
	2, 38, 162, 3, 2, 2, 2, 40, 175, 3, 2, 2, 2, 42, 183, 3, 2, 2, 2, 44, 46,
	5, 4, 3, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2,
	47, 48, 3, 2, 2, 2, 48, 3, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 53, 5, 6,
	4, 2, 51, 53, 5, 8, 5, 2, 52, 50, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 5,
	3, 2, 2, 2, 54, 55, 7, 3, 2, 2, 55, 56, 7, 24, 2, 2, 56, 60, 7, 25, 2,
	2, 57, 59, 5, 10, 6, 2, 58, 57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58,
	3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2,
	63, 64, 7, 4, 2, 2, 64, 7, 3, 2, 2, 2, 65, 69, 7, 5, 2, 2, 66, 68, 5, 10,
	6, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70,
	3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 6, 2, 2,
	73, 9, 3, 2, 2, 2, 74, 92, 5, 6, 4, 2, 75, 92, 5, 8, 5, 2, 76, 92, 5, 14,
	8, 2, 77, 92, 5, 16, 9, 2, 78, 92, 5, 18, 10, 2, 79, 92, 5, 20, 11, 2,
	80, 92, 5, 22, 12, 2, 81, 92, 5, 24, 13, 2, 82, 92, 5, 26, 14, 2, 83, 92,
	5, 28, 15, 2, 84, 92, 5, 30, 16, 2, 85, 92, 5, 32, 17, 2, 86, 92, 5, 38,
	20, 2, 87, 92, 5, 40, 21, 2, 88, 92, 5, 42, 22, 2, 89, 92, 5, 34, 18, 2,
	90, 92, 5, 36, 19, 2, 91, 74, 3, 2, 2, 2, 91, 75, 3, 2, 2, 2, 91, 76, 3,
	2, 2, 2, 91, 77, 3, 2, 2, 2, 91, 78, 3, 2, 2, 2, 91, 79, 3, 2, 2, 2, 91,
	80, 3, 2, 2, 2, 91, 81, 3, 2, 2, 2, 91, 82, 3, 2, 2, 2, 91, 83, 3, 2, 2,
	2, 91, 84, 3, 2, 2, 2, 91, 85, 3, 2, 2, 2, 91, 86, 3, 2, 2, 2, 91, 87,
	3, 2, 2, 2, 91, 88, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 90, 3, 2, 2, 2,
	92, 11, 3, 2, 2, 2, 93, 103, 5, 22, 12, 2, 94, 103, 5, 24, 13, 2, 95, 103,
	5, 26, 14, 2, 96, 103, 5, 28, 15, 2, 97, 103, 5, 30, 16, 2, 98, 103, 5,
	14, 8, 2, 99, 103, 5, 16, 9, 2, 100, 103, 5, 36, 19, 2, 101, 103, 5, 32,
	17, 2, 102, 93, 3, 2, 2, 2, 102, 94, 3, 2, 2, 2, 102, 95, 3, 2, 2, 2, 102,
	96, 3, 2, 2, 2, 102, 97, 3, 2, 2, 2, 102, 98, 3, 2, 2, 2, 102, 99, 3, 2,
	2, 2, 102, 100, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 13, 3, 2, 2, 2,
	104, 108, 9, 2, 2, 2, 105, 106, 7, 7, 2, 2, 106, 107, 9, 3, 2, 2, 107,
	109, 7, 6, 2, 2, 108, 105, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 15, 3,
	2, 2, 2, 110, 114, 7, 21, 2, 2, 111, 112, 7, 7, 2, 2, 112, 113, 9, 3, 2,
	2, 113, 115, 7, 6, 2, 2, 114, 111, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115,
	17, 3, 2, 2, 2, 116, 117, 7, 8, 2, 2, 117, 118, 9, 2, 2, 2, 118, 19, 3,
	2, 2, 2, 119, 120, 7, 8, 2, 2, 120, 121, 7, 21, 2, 2, 121, 21, 3, 2, 2,
	2, 122, 126, 9, 4, 2, 2, 123, 124, 7, 7, 2, 2, 124, 125, 9, 2, 2, 2, 125,
	127, 7, 6, 2, 2, 126, 123, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 23, 3,
	2, 2, 2, 128, 132, 7, 13, 2, 2, 129, 130, 7, 7, 2, 2, 130, 131, 9, 2, 2,
	2, 131, 133, 7, 6, 2, 2, 132, 129, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133,
	25, 3, 2, 2, 2, 134, 138, 7, 15, 2, 2, 135, 136, 7, 7, 2, 2, 136, 137,
	9, 2, 2, 2, 137, 139, 7, 6, 2, 2, 138, 135, 3, 2, 2, 2, 138, 139, 3, 2,
	2, 2, 139, 27, 3, 2, 2, 2, 140, 144, 7, 16, 2, 2, 141, 142, 7, 7, 2, 2,
	142, 143, 9, 2, 2, 2, 143, 145, 7, 6, 2, 2, 144, 141, 3, 2, 2, 2, 144,
	145, 3, 2, 2, 2, 145, 29, 3, 2, 2, 2, 146, 150, 7, 17, 2, 2, 147, 148,
	7, 7, 2, 2, 148, 149, 9, 2, 2, 2, 149, 151, 7, 6, 2, 2, 150, 147, 3, 2,
	2, 2, 150, 151, 3, 2, 2, 2, 151, 31, 3, 2, 2, 2, 152, 156, 7, 27, 2, 2,
	153, 154, 7, 7, 2, 2, 154, 155, 9, 2, 2, 2, 155, 157, 7, 6, 2, 2, 156,
	153, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 33, 3, 2, 2, 2, 158, 159, 9,
	5, 2, 2, 159, 35, 3, 2, 2, 2, 160, 161, 7, 28, 2, 2, 161, 37, 3, 2, 2,
	2, 162, 166, 7, 9, 2, 2, 163, 165, 5, 12, 7, 2, 164, 163, 3, 2, 2, 2, 165,
	168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 169,
	3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 173, 7, 6, 2, 2, 170, 171, 7, 7,
	2, 2, 171, 172, 9, 3, 2, 2, 172, 174, 7, 6, 2, 2, 173, 170, 3, 2, 2, 2,
	173, 174, 3, 2, 2, 2, 174, 39, 3, 2, 2, 2, 175, 177, 7, 10, 2, 2, 176,
	178, 5, 12, 7, 2, 177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 177,
	3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 7, 6,
	2, 2, 182, 41, 3, 2, 2, 2, 183, 187, 9, 6, 2, 2, 184, 186, 5, 10, 6, 2,
	185, 184, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 190, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 191,
	7, 6, 2, 2, 191, 43, 3, 2, 2, 2, 20, 47, 52, 60, 69, 91, 102, 108, 114,
	126, 132, 138, 144, 150, 156, 166, 173, 179, 187,
}
var literalNames = []string{
	"", "'['", "';;'", "'('", "')'", "'.('", "'`'", "'(*'", "'(?'", "'(true'",
	"'(false'", "", "", "", "", "", "", "", "", "", "", "'/'", "", "':'", "';'",
	"", "'|'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER",
	"FLOAT_NUMBER", "STRING", "COMPLEX_NUMBER", "TRUE", "FALSE", "SYS", "CMD",
	"SYSF", "SLASH", "NAME", "TOBEGIN", "TOEND", "GLOB", "SEPARATE", "COMMENT",
	"BLOCK_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"expressions", "root_term", "ns", "block", "term", "data", "call_term",
	"operator_term", "ref_call_term", "ref_operator_term", "boolean_term",
	"integer_term", "float_term", "string_term", "complex_term", "glob_term",
	"mode_term", "separate_term", "datablock_term", "matchblock_term", "logicblock_term",
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
	BundParserT__7            = 8
	BundParserT__8            = 9
	BundParserT__9            = 10
	BundParserINTEGER         = 11
	BundParserDECIMAL_INTEGER = 12
	BundParserFLOAT_NUMBER    = 13
	BundParserSTRING          = 14
	BundParserCOMPLEX_NUMBER  = 15
	BundParserTRUE            = 16
	BundParserFALSE           = 17
	BundParserSYS             = 18
	BundParserCMD             = 19
	BundParserSYSF            = 20
	BundParserSLASH           = 21
	BundParserNAME            = 22
	BundParserTOBEGIN         = 23
	BundParserTOEND           = 24
	BundParserGLOB            = 25
	BundParserSEPARATE        = 26
	BundParserCOMMENT         = 27
	BundParserBLOCK_COMMENT   = 28
	BundParserWS              = 29
	BundParserSHEBANG         = 30
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
	BundParserRULE_matchblock_term   = 19
	BundParserRULE_logicblock_term   = 20
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserT__0 || _la == BundParserT__2 {
		{
			p.SetState(42)
			p.Root_term()
		}

		p.SetState(47)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(48)
			p.Ns()
		}

	case BundParserT__2:
		{
			p.SetState(49)
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
		p.SetState(52)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(53)

		var _m = p.Match(BundParserNAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(54)
		p.Match(BundParserTOBEGIN)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(55)

			var _x = p.Term()

			localctx.(*NsContext)._term = _x
		}
		localctx.(*NsContext).body = append(localctx.(*NsContext).body, localctx.(*NsContext)._term)

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
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
		p.SetState(63)
		p.Match(BundParserT__2)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(64)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).body = append(localctx.(*BlockContext).body, localctx.(*BlockContext)._term)

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
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

func (s *TermContext) Matchblock_term() IMatchblock_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchblock_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchblock_termContext)
}

func (s *TermContext) Logicblock_term() ILogicblock_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicblock_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicblock_termContext)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(72)
			p.Ns()
		}

	case 2:
		{
			p.SetState(73)
			p.Block()
		}

	case 3:
		{
			p.SetState(74)
			p.Call_term()
		}

	case 4:
		{
			p.SetState(75)
			p.Operator_term()
		}

	case 5:
		{
			p.SetState(76)
			p.Ref_call_term()
		}

	case 6:
		{
			p.SetState(77)
			p.Ref_operator_term()
		}

	case 7:
		{
			p.SetState(78)
			p.Boolean_term()
		}

	case 8:
		{
			p.SetState(79)
			p.Integer_term()
		}

	case 9:
		{
			p.SetState(80)
			p.Float_term()
		}

	case 10:
		{
			p.SetState(81)
			p.String_term()
		}

	case 11:
		{
			p.SetState(82)
			p.Complex_term()
		}

	case 12:
		{
			p.SetState(83)
			p.Glob_term()
		}

	case 13:
		{
			p.SetState(84)
			p.Datablock_term()
		}

	case 14:
		{
			p.SetState(85)
			p.Matchblock_term()
		}

	case 15:
		{
			p.SetState(86)
			p.Logicblock_term()
		}

	case 16:
		{
			p.SetState(87)
			p.Mode_term()
		}

	case 17:
		{
			p.SetState(88)
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

func (s *DataContext) Glob_term() IGlob_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlob_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlob_termContext)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserTRUE, BundParserFALSE:
		{
			p.SetState(91)
			p.Boolean_term()
		}

	case BundParserINTEGER:
		{
			p.SetState(92)
			p.Integer_term()
		}

	case BundParserFLOAT_NUMBER:
		{
			p.SetState(93)
			p.Float_term()
		}

	case BundParserSTRING:
		{
			p.SetState(94)
			p.String_term()
		}

	case BundParserCOMPLEX_NUMBER:
		{
			p.SetState(95)
			p.Complex_term()
		}

	case BundParserSYSF, BundParserNAME:
		{
			p.SetState(96)
			p.Call_term()
		}

	case BundParserCMD:
		{
			p.SetState(97)
			p.Operator_term()
		}

	case BundParserSEPARATE:
		{
			p.SetState(98)
			p.Separate_term()
		}

	case BundParserGLOB:
		{
			p.SetState(99)
			p.Glob_term()
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
		p.SetState(102)

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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(103)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(104)

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
			p.SetState(105)
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
		p.SetState(108)

		var _m = p.Match(BundParserCMD)

		localctx.(*Operator_termContext).VALUE = _m
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(109)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(110)

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
			p.SetState(111)
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
		p.SetState(114)
		p.Match(BundParserT__5)
	}
	{
		p.SetState(115)

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
		p.SetState(117)
		p.Match(BundParserT__5)
	}
	{
		p.SetState(118)

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
		p.SetState(120)

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
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(121)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(122)

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
			p.SetState(123)
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
		p.SetState(126)

		var _m = p.Match(BundParserINTEGER)

		localctx.(*Integer_termContext).VALUE = _m
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(127)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(128)

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
			p.SetState(129)
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
		p.SetState(132)

		var _m = p.Match(BundParserFLOAT_NUMBER)

		localctx.(*Float_termContext).VALUE = _m
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(133)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(134)

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
			p.SetState(135)
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
		p.SetState(138)

		var _m = p.Match(BundParserSTRING)

		localctx.(*String_termContext).VALUE = _m
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(139)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(140)

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
			p.SetState(141)
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
		p.SetState(144)

		var _m = p.Match(BundParserCOMPLEX_NUMBER)

		localctx.(*Complex_termContext).VALUE = _m
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(145)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(146)

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
			p.SetState(147)
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
		p.SetState(150)

		var _m = p.Match(BundParserGLOB)

		localctx.(*Glob_termContext).VALUE = _m
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(151)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(152)

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
			p.SetState(153)
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
		p.SetState(156)

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
		p.SetState(158)

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
		p.SetState(160)
		p.Match(BundParserT__6)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(161)

			var _x = p.Data()

			localctx.(*Datablock_termContext)._data = _x
		}
		localctx.(*Datablock_termContext).body = append(localctx.(*Datablock_termContext).body, localctx.(*Datablock_termContext)._data)

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(BundParserT__3)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserT__4 {
		{
			p.SetState(168)
			p.Match(BundParserT__4)
		}
		{
			p.SetState(169)

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
			p.SetState(170)
			p.Match(BundParserT__3)
		}

	}

	return localctx
}

// IMatchblock_termContext is an interface to support dynamic dispatch.
type IMatchblock_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_data returns the _data rule contexts.
	Get_data() IDataContext

	// Set_data sets the _data rule contexts.
	Set_data(IDataContext)

	// GetBody returns the body rule context list.
	GetBody() []IDataContext

	// SetBody sets the body rule context list.
	SetBody([]IDataContext)

	// IsMatchblock_termContext differentiates from other interfaces.
	IsMatchblock_termContext()
}

type Matchblock_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_data  IDataContext
	body   []IDataContext
}

func NewEmptyMatchblock_termContext() *Matchblock_termContext {
	var p = new(Matchblock_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_matchblock_term
	return p
}

func (*Matchblock_termContext) IsMatchblock_termContext() {}

func NewMatchblock_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matchblock_termContext {
	var p = new(Matchblock_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_matchblock_term

	return p
}

func (s *Matchblock_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Matchblock_termContext) Get_data() IDataContext { return s._data }

func (s *Matchblock_termContext) Set_data(v IDataContext) { s._data = v }

func (s *Matchblock_termContext) GetBody() []IDataContext { return s.body }

func (s *Matchblock_termContext) SetBody(v []IDataContext) { s.body = v }

func (s *Matchblock_termContext) AllData() []IDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataContext)(nil)).Elem())
	var tst = make([]IDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataContext)
		}
	}

	return tst
}

func (s *Matchblock_termContext) Data(i int) IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *Matchblock_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matchblock_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matchblock_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterMatchblock_term(s)
	}
}

func (s *Matchblock_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitMatchblock_term(s)
	}
}

func (p *BundParser) Matchblock_term() (localctx IMatchblock_termContext) {
	localctx = NewMatchblock_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BundParserRULE_matchblock_term)
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
		p.SetState(173)
		p.Match(BundParserT__7)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0) {
		{
			p.SetState(174)

			var _x = p.Data()

			localctx.(*Matchblock_termContext)._data = _x
		}
		localctx.(*Matchblock_termContext).body = append(localctx.(*Matchblock_termContext).body, localctx.(*Matchblock_termContext)._data)

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(179)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ILogicblock_termContext is an interface to support dynamic dispatch.
type ILogicblock_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHDR returns the HDR token.
	GetHDR() antlr.Token

	// SetHDR sets the HDR token.
	SetHDR(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsLogicblock_termContext differentiates from other interfaces.
	IsLogicblock_termContext()
}

type Logicblock_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	HDR    antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLogicblock_termContext() *Logicblock_termContext {
	var p = new(Logicblock_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_logicblock_term
	return p
}

func (*Logicblock_termContext) IsLogicblock_termContext() {}

func NewLogicblock_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logicblock_termContext {
	var p = new(Logicblock_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_logicblock_term

	return p
}

func (s *Logicblock_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Logicblock_termContext) GetHDR() antlr.Token { return s.HDR }

func (s *Logicblock_termContext) SetHDR(v antlr.Token) { s.HDR = v }

func (s *Logicblock_termContext) Get_term() ITermContext { return s._term }

func (s *Logicblock_termContext) Set_term(v ITermContext) { s._term = v }

func (s *Logicblock_termContext) GetBody() []ITermContext { return s.body }

func (s *Logicblock_termContext) SetBody(v []ITermContext) { s.body = v }

func (s *Logicblock_termContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Logicblock_termContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Logicblock_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logicblock_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logicblock_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLogicblock_term(s)
	}
}

func (s *Logicblock_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLogicblock_term(s)
	}
}

func (p *BundParser) Logicblock_term() (localctx ILogicblock_termContext) {
	localctx = NewLogicblock_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BundParserRULE_logicblock_term)
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
		p.SetState(181)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Logicblock_termContext).HDR = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserT__8 || _la == BundParserT__9) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Logicblock_termContext).HDR = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserCMD)|(1<<BundParserSYSF)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserGLOB)|(1<<BundParserSEPARATE))) != 0 {
		{
			p.SetState(182)

			var _x = p.Term()

			localctx.(*Logicblock_termContext)._term = _x
		}
		localctx.(*Logicblock_termContext).body = append(localctx.(*Logicblock_termContext).body, localctx.(*Logicblock_termContext)._term)

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(188)
		p.Match(BundParserT__3)
	}

	return localctx
}
