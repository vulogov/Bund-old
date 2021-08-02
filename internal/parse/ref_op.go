package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterRef_operator_term(c *parser.Ref_operator_termContext) {
	l.VM.Opcode("ROP").InParser(l.VM, c.GetVALUE().GetText())
}

func (l *bundExecListener) EnterRef_operator_term(c *parser.Ref_operator_termContext) {
	l.VM.RunOp("ROP", c.GetVALUE().GetText())
}
