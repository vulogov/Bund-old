package parse

import (
	"strings"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterString_term(c *parser.String_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	l.VM.Opcode("str").InParser(l.VM, c.GetVALUE().GetText(), functor)
}

func (l *bundExecListener) EnterString_term(c *parser.String_termContext) {
	var functor string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	tstring := c.GetVALUE().GetText()
	if strings.HasPrefix(tstring, "\"\"\"") {
		tstring = strings.TrimPrefix(tstring, "\"\"\"")
	} else if strings.HasPrefix(tstring, "\"") {
		tstring = strings.TrimPrefix(tstring, "\"")
	} else if strings.HasPrefix(tstring, "'''") {
		tstring = strings.TrimPrefix(tstring, "'''")
	} else if strings.HasPrefix(tstring, "'") {
		tstring = strings.TrimPrefix(tstring, "'")
	}
	if strings.HasSuffix(tstring, "\"\"\"") {
		tstring = strings.TrimSuffix(tstring, "\"\"\"")
	} else if strings.HasSuffix(tstring, "\"") {
		tstring = strings.TrimSuffix(tstring, "\"")
	} else if strings.HasSuffix(tstring, "'''") {
		tstring = strings.TrimSuffix(tstring, "'''")
	} else if strings.HasSuffix(tstring, "'") {
		tstring = strings.TrimSuffix(tstring, "'")
	}
	l.VM.RunOp("str", tstring, functor)
}
