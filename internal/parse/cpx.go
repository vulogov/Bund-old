package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterComplex_term(c *parser.Complex_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("cpx").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterComplex_term(c *parser.Complex_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("cpx", c.GetVALUE().GetText(), functor)
}
