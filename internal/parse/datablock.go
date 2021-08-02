package parse

import (
	"github.com/google/uuid"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterDatablock_term(c *parser.Datablock_termContext) {
	l.VM.Opcode("DBLOCK").InParser(l.VM)
}

func (l *bundListener) ExitDatablock_term(c *parser.Datablock_termContext) {
	l.VM.Opcode("exitDBLOCK").InParser(l.VM)
}

func (l *bundExecListener) EnterDatablock_term(c *parser.Datablock_termContext) {
	blockname := uuid.New().String()
	l.VM.RunOp("DBLOCK", blockname)
}

func (l *bundExecListener) ExitDatablock_term(c *parser.Datablock_termContext) {
	l.VM.RunOp("exitDBLOCK")
}
