package bund

import (
	"github.com/cosiner/argv"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/conf"
	tlog "github.com/vulogov/Bund/internal/log"
	"github.com/vulogov/Bund/internal/signal"
)

func Init() {
	tlog.Init()
	log.Debug("[ BUND ] bund.Init() is reached")
	signal.InitSignal()
	if len(*conf.Args) > 0 {
		Argv, err := argv.Argv(*conf.Args, func(backquoted string) (string, error) {
			return backquoted, nil
		}, nil)
		if err != nil {
			log.Fatalf("Error parsing ARGS: %v", err)
		}
		log.Debugf("ARGV: %v", Argv)
		conf.Argv = Argv
	}
}
