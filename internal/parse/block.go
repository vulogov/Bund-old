package parse

import (
	"github.com/google/uuid"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterBlock(c *parser.BlockContext) {
	l.VM.Opcode("BLOCK").InParser(l.VM)
}

func (l *bundListener) ExitBlock(c *parser.BlockContext) {
	l.VM.Opcode("exitBLOCK").InParser(l.VM)
}

func (l *bundExecListener) EnterBlock(c *parser.BlockContext) {
	blockname := uuid.New().String()
	l.VM.RunOp("BLOCK", blockname)
}

func (l *bundExecListener) ExitBlock(c *parser.BlockContext) {
	l.VM.RunOp("exitBLOCK")
}
