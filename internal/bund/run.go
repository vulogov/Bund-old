package bund

import (
	"io/ioutil"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/parse"
	"github.com/vulogov/Bund/internal/signal"
)

func Run() {
	Init()
	log.Debug("[ BUND ] tsak.Run() is reached")
	for _, s := range *conf.Scripts {
		log.Debugf("[ BUND ] Loading: %v", s)
		code, err := ioutil.ReadFile(s)
		if err != nil {
			log.Errorf("Error loading %v: %v", s, err)
			continue
		}
		parse.ParserExec(s, string(code))
	}
	if *conf.Main != "" {
		log.Debugf("[ BUND ] Executing main script: %v", *conf.Main)
		code, err := ioutil.ReadFile(*conf.Main)
		if err != nil {
			log.Errorf("Error loading %v: %v", *conf.Main, err)
		} else {
			parse.ParserExec(*conf.Main, string(code))
		}
	}
	signal.ExitRequest()
}
