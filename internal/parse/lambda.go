package parse

import (
	"github.com/google/uuid"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterLambda_term(c *parser.Lambda_termContext) {
	blockname := uuid.New().String()
	l.VM.Opcode("LAMBDA").InParser(l.VM, blockname)
}

func (l *bundListener) ExitLambda_term(c *parser.Lambda_termContext) {
	l.VM.Opcode("exitLAMBDA").InParser(l.VM)
}

func (l *bundExecListener) EnterLambda_term(c *parser.Lambda_termContext) {
	blockname := uuid.New().String()
	l.VM.LambdaRunOp("LAMBDA", blockname)
}

func (l *bundExecListener) ExitLambda_term(c *parser.Lambda_termContext) {
	l.VM.LambdaRunOp("exitLAMBDA")
}
