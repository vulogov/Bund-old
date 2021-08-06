package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterThing_term(c *parser.Thing_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("flt").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterThing_term(c *parser.Thing_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("flt", c.GetVALUE().GetText(), functor)
}
