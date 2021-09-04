package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterRef_call_term(c *parser.Ref_call_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("RCALL").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterRef_call_term(c *parser.Ref_call_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("RCALL", c.GetVALUE().GetText(), functor)
}
