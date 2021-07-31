package parse

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
	"github.com/vulogov/Bund/internal/vm"
)

type bundListener struct {
	*parser.BaseBundListener
	VM     *vm.VM
	errors int
}

type bundErrorListener struct {
	antlr.ErrorListener
	VM     *vm.VM
	errors int
}

func ParserPrint(code string) {
	errorListener := new(bundErrorListener)
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	listener := new(bundListener)
	listener.VM = vm.NewVM("<parser>")
	errorListener.VM = listener.VM
	listener.VM.Debug("Code passed on parser print: %v", code)
	if errorListener.errors > 0 {
		listener.VM.Error("%v lexer errors detected.", errorListener.errors)
		return
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
	if errorListener.errors > 0 {
		listener.VM.Error("Errors detected: %v", errorListener.errors)
	} else {
		listener.VM.Debug("No errors detected")
	}
}

func (l *bundErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.VM.Error("Syntax error line=%v, column=%v : %v", line, column, msg)
	l.errors += 1
}
func (l *bundErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.VM.Error("Ambiguity Error")
	l.errors += 1
}
func (l *bundErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.VM.Error("Attempting in Full Context")
	l.errors += 1
}
func (l *bundErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	l.VM.Error("Context sensitivity error")
	l.errors += 1
}
