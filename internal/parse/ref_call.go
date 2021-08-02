package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterRef_call_term(c *parser.Ref_call_termContext) {
	l.VM.Opcode("RCALL").InParser(l.VM, c.GetVALUE().GetText())
}

func (l *bundExecListener) EnterRef_call_term(c *parser.Ref_call_termContext) {
	l.VM.RunOp("RCALL", c.GetVALUE().GetText())
}
