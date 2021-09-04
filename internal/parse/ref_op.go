package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterRef_operator_term(c *parser.Ref_operator_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("ROP").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterRef_operator_term(c *parser.Ref_operator_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("ROP", c.GetVALUE().GetText(), functor)
}
