package parse

import (
	"strings"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterGlob_term(c *parser.Glob_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("glob").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterGlob_term(c *parser.Glob_termContext) {
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
	l.VM.RunOp("glob", tstring, functor)
}
