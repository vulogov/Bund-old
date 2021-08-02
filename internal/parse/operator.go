package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterOperator_term(c *parser.Operator_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("OP").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterOperator_term(c *parser.Operator_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.RunOp("OP", c.GetVALUE().GetText(), functor)
}
