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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 103,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 45, 10, 8, 12, 8, 14,
	8, 48, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 54, 10, 9, 12, 9, 14, 9, 57,
	11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 65, 10, 10, 12, 10,
	14, 10, 68, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 6, 11, 76,
	10, 11, 13, 11, 14, 11, 77, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 85,
	10, 12, 12, 12, 14, 12, 88, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13,
	94, 10, 13, 3, 13, 5, 13, 97, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14, 102,
	10, 14, 3, 66, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 2, 27, 2, 3, 2, 7, 4, 2, 12, 12, 15, 15,
	5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 92, 99, 124, 3, 2, 99, 124, 3,
	2, 50, 59, 2, 109, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 36, 3,
	2, 2, 2, 11, 38, 3, 2, 2, 2, 13, 40, 3, 2, 2, 2, 15, 42, 3, 2, 2, 2, 17,
	49, 3, 2, 2, 2, 19, 60, 3, 2, 2, 2, 21, 75, 3, 2, 2, 2, 23, 81, 3, 2, 2,
	2, 25, 96, 3, 2, 2, 2, 27, 101, 3, 2, 2, 2, 29, 30, 7, 93, 2, 2, 30, 4,
	3, 2, 2, 2, 31, 32, 7, 60, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 61, 2, 2,
	34, 35, 7, 61, 2, 2, 35, 8, 3, 2, 2, 2, 36, 37, 7, 42, 2, 2, 37, 10, 3,
	2, 2, 2, 38, 39, 7, 43, 2, 2, 39, 12, 3, 2, 2, 2, 40, 41, 7, 49, 2, 2,
	41, 14, 3, 2, 2, 2, 42, 46, 5, 25, 13, 2, 43, 45, 5, 27, 14, 2, 44, 43,
	3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2,
	47, 16, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 37, 2, 2, 50, 51, 7,
	37, 2, 2, 51, 55, 3, 2, 2, 2, 52, 54, 10, 2, 2, 2, 53, 52, 3, 2, 2, 2,
	54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3,
	2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 8, 9, 2, 2, 59, 18, 3, 2, 2, 2, 60,
	61, 7, 49, 2, 2, 61, 62, 7, 44, 2, 2, 62, 66, 3, 2, 2, 2, 63, 65, 11, 2,
	2, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 66, 64,
	3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7, 44, 2, 2,
	70, 71, 7, 49, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 8, 10, 2, 2, 73, 20, 3,
	2, 2, 2, 74, 76, 9, 3, 2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 8, 11,
	2, 2, 80, 22, 3, 2, 2, 2, 81, 82, 7, 37, 2, 2, 82, 86, 7, 35, 2, 2, 83,
	85, 10, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2,
	2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90,
	8, 12, 3, 2, 90, 24, 3, 2, 2, 2, 91, 94, 9, 4, 2, 2, 92, 94, 5, 13, 7,
	2, 93, 91, 3, 2, 2, 2, 93, 92, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 97,
	9, 5, 2, 2, 96, 93, 3, 2, 2, 2, 96, 95, 3, 2, 2, 2, 97, 26, 3, 2, 2, 2,
	98, 102, 5, 25, 13, 2, 99, 102, 9, 6, 2, 2, 100, 102, 5, 13, 7, 2, 101,
	98, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 28, 3,
	2, 2, 2, 11, 2, 46, 55, 66, 77, 86, 93, 96, 101, 4, 8, 2, 2, 2, 3, 2,
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
	"", "", "", "", "", "", "SLASH", "NAME", "COMMENT", "BLOCK_COMMENT", "WS",
	"SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "SLASH", "NAME", "COMMENT", "BLOCK_COMMENT",
	"WS", "SHEBANG", "ID_START", "ID_CONTINUE",
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
	BundLexerT__0          = 1
	BundLexerT__1          = 2
	BundLexerT__2          = 3
	BundLexerT__3          = 4
	BundLexerT__4          = 5
	BundLexerSLASH         = 6
	BundLexerNAME          = 7
	BundLexerCOMMENT       = 8
	BundLexerBLOCK_COMMENT = 9
	BundLexerWS            = 10
	BundLexerSHEBANG       = 11
)
