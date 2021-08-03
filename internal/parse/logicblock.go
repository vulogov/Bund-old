package parse

import (
	"github.com/cstockton/go-conv"
	"github.com/google/uuid"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterLogicblock_term(c *parser.Logicblock_termContext) {
	blockname := uuid.New().String()
	v, _ := conv.Bool(c.GetHDR().GetText()[1:])
	l.VM.Opcode("LBLOCK").InParser(l.VM, blockname, "", v)
}

func (l *bundListener) ExitLogicblock_term(c *parser.Logicblock_termContext) {
	v, _ := conv.Bool(c.GetHDR().GetText()[1:])
	l.VM.Opcode("exitLBLOCK").InParser(l.VM, v)
}

func (l *bundExecListener) EnterLogicblock_term(c *parser.Logicblock_termContext) {
	blockname := uuid.New().String()
	v, _ := conv.Bool(c.GetHDR().GetText()[1:])
	l.VM.AlwaysRunOp("LBLOCK", blockname, "", v)
}

func (l *bundExecListener) ExitLogicblock_term(c *parser.Logicblock_termContext) {
	v, _ := conv.Bool(c.GetHDR().GetText()[1:])
	l.VM.AlwaysRunOp("exitLBLOCK", v)
}
