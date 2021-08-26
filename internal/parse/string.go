package parse

import (
	"strings"

	"github.com/vulogov/Bund/internal/parser"
)

func (l *bundListener) EnterString_term(c *parser.String_termContext) {
	var functor string
	var prefunction string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	if c.GetPRE() != nil {
		prefunction = c.GetPRE().GetText()
	}
	l.VM.Opcode("str").InParser(l.VM, c.GetVALUE().GetText(), functor, prefunction)
}

func (l *bundExecListener) EnterString_term(c *parser.String_termContext) {
	var functor string
	var prefunction string
	if c.GetFUNCTOR() != nil {
		functor = c.GetFUNCTOR().GetText()
	}
	if c.GetPRE() != nil {
		prefunction = c.GetPRE().GetText()
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
	if c.GetPRE() != nil {
		l.VM.Debug("Executing prefunction %v for the string", prefunction)
		_, err := l.VM.GetType(prefunction)
		if err == nil {
			l.VM.Debug("Datatype %v is found, convert string representation to that type", prefunction)
			l.VM.RunOp(prefunction, tstring, functor)
			return
		}
		f := l.VM.GetFunction(prefunction)
		if f != nil {
			l.VM.Debug("Embedded Function %v is found. Pass string to that function", prefunction)
			eh, err := l.VM.GetType("str")
			l.VM.OnError(err, "Error in string prefunction compute")
			v := eh.Factory(l.VM, tstring)
			v.Functor = functor
			res, err := f(l.VM, v)
			l.VM.OnError(err, "Error calling %v", prefunction)
			if res != nil {
				l.VM.EvalCmd(res)
			} else {
				l.VM.Warning("Prefunction %v returned NULL", prefunction)
			}
		}
		l.VM.Debug("Prefunction %v not found, continue as string", prefunction)
		l.VM.RunOp("str", tstring, functor)
	} else {
		l.VM.RunOp("str", tstring, functor)
	}
}
