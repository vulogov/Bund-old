package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterCall_term(c *parser.Call_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("CALL").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterCall_term(c *parser.Call_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("CALL", c.GetVALUE().GetText(), functor)
}
