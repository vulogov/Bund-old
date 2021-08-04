package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterFunction_term(c *parser.Function_termContext) {
	l.VM.Opcode("FUNCTION").InParser(l.VM, c.GetName().GetText())
}

func (l *bundListener) ExitFunction_term(c *parser.Function_termContext) {
	l.VM.Opcode("exitFUNCTION").InParser(l.VM, c.GetName().GetText())
}

func (l *bundExecListener) EnterFunction_term(c *parser.Function_termContext) {
	l.VM.LambdaRunOp("FUNCTION", c.GetName().GetText())
}

func (l *bundExecListener) ExitFunction_term(c *parser.Function_termContext) {
	l.VM.LambdaRunOp("exitFUNCTION", c.GetName().GetText())
}
