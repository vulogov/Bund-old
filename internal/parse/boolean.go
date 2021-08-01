package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterBoolean_term(c *parser.Boolean_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("bool").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterBoolean_term(c *parser.Boolean_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("bool", c.GetVALUE().GetText(), functor)
}
