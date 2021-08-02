package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterMode_term(c *parser.Mode_termContext) {
	l.VM.Opcode("MODE").InParser(l.VM, c.GetVALUE().GetText())
}

func (l *bundExecListener) EnterMode_term(c *parser.Mode_termContext) {
	l.VM.RunOp("MODE", c.GetVALUE().GetText())
}
