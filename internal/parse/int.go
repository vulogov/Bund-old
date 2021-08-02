package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterInteger_term(c *parser.Integer_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("int").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterInteger_term(c *parser.Integer_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("int", c.GetVALUE().GetText(), functor)
}
