package bund

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/parse"
	"github.com/vulogov/Bund/internal/signal"
)

func Eval() {
	Init()
	log.Debug("[ BUND ] tsak.Eval() is reached")
	if *conf.LexerPrint {
		parse.LexerPrint(*conf.Expr)
		return
	}
	if *conf.ParserPrint {
		parse.ParserPrint(*conf.Expr)
		return
	}
	parse.ParserExec(nil, "<eval>", *conf.Expr)
	signal.ExitRequest()
}
