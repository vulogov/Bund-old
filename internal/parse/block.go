package parse

import (
	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterBlock(c *parser.BlockContext) {
	l.VM.Opcode("BLOCK").InParser(l.VM)
}

func (l *bundListener) ExitBlock(c *parser.BlockContext) {
	l.VM.Opcode("exitBLOCK").InParser(l.VM)
}

func (l *bundExecListener) EnterBlock(c *parser.BlockContext) {

}

func (l *bundExecListener) ExitBlock(c *parser.BlockContext) {

}
