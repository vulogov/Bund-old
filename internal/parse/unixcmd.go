package parse

import (
	"strings"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterUnixcmd_term(c *parser.Unixcmd_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("unixcmd").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterUnixcmd_term(c *parser.Unixcmd_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	tstring := c.GetVALUE().GetText()[1:]
	if strings.HasPrefix(tstring, "\"") {
		tstring = strings.TrimPrefix(tstring, "\"")
	} else if strings.HasPrefix(tstring, "'") {
		tstring = strings.TrimPrefix(tstring, "'")
	}
	if strings.HasSuffix(tstring, "\"") {
		tstring = strings.TrimSuffix(tstring, "\"")
	} else if strings.HasSuffix(tstring, "'") {
		tstring = strings.TrimSuffix(tstring, "'")
	}
	l.VM.RunOp("unixcmd", tstring, functor)
}
