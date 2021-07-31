package parse

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
)

func LexerPrint(code string) {
	var n = 0
	var group = true
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		if n < 0 {
			n = 0
		}
		ident := strings.Repeat("..", n)
		fmt.Printf("%s %s (%q)\n", ident,
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		if t.GetText() == "|" {
			if group {
				n++
			} else {
				n--
			}
			group = !group
		}
		if t.GetText() == "(" {
			n++
		}
		if t.GetText() == "[" {
			n++
		}
		if t.GetText() == ")" {
			n--
		}
		if t.GetText() == "]" {
			n--
		}
		if t.GetText() == ";;" {
			n--
		}
	}
}
