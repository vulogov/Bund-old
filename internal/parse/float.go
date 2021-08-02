package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterFloat_term(c *parser.Float_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("flt").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterFloat_term(c *parser.Float_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("flt", c.GetVALUE().GetText(), functor)
}
