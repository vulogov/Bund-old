// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 53, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 37, 10,
	8, 12, 8, 14, 8, 40, 11, 8, 3, 9, 3, 9, 5, 9, 44, 10, 9, 3, 9, 5, 9, 47,
	10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 52, 10, 10, 2, 2, 11, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 2, 19, 2, 3, 2, 5, 4, 2, 67, 92, 99,
	124, 3, 2, 99, 124, 3, 2, 50, 59, 2, 55, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25,
	3, 2, 2, 2, 9, 28, 3, 2, 2, 2, 11, 30, 3, 2, 2, 2, 13, 32, 3, 2, 2, 2,
	15, 34, 3, 2, 2, 2, 17, 46, 3, 2, 2, 2, 19, 51, 3, 2, 2, 2, 21, 22, 7,
	93, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 60, 2, 2, 24, 6, 3, 2, 2, 2, 25,
	26, 7, 61, 2, 2, 26, 27, 7, 61, 2, 2, 27, 8, 3, 2, 2, 2, 28, 29, 7, 42,
	2, 2, 29, 10, 3, 2, 2, 2, 30, 31, 7, 43, 2, 2, 31, 12, 3, 2, 2, 2, 32,
	33, 7, 49, 2, 2, 33, 14, 3, 2, 2, 2, 34, 38, 5, 17, 9, 2, 35, 37, 5, 19,
	10, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38,
	39, 3, 2, 2, 2, 39, 16, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 44, 9, 2, 2,
	2, 42, 44, 5, 13, 7, 2, 43, 41, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 47,
	3, 2, 2, 2, 45, 47, 9, 3, 2, 2, 46, 43, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2,
	47, 18, 3, 2, 2, 2, 48, 52, 5, 17, 9, 2, 49, 52, 9, 4, 2, 2, 50, 52, 5,
	13, 7, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	20, 3, 2, 2, 2, 7, 2, 38, 43, 46, 51, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "':'", "';;'", "'('", "')'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "SLASH", "NAME",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "SLASH", "NAME", "ID_START", "ID_CONTINUE",
}

type BundLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewBundLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *BundLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundLexer(input antlr.CharStream) *BundLexer {
	l := new(BundLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bund.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BundLexer tokens.
const (
	BundLexerT__0  = 1
	BundLexerT__1  = 2
	BundLexerT__2  = 3
	BundLexerT__3  = 4
	BundLexerT__4  = 5
	BundLexerSLASH = 6
	BundLexerNAME  = 7
)
