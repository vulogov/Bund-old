package bund

import (
	"github.com/pieterclaerhout/go-log"

	tlog "github.com/vulogov/Bund/internal/log"
	"github.com/vulogov/Bund/internal/signal"
)

func Init() {
	tlog.Init()
	log.Debug("[ BUND ] bund.Init() is reached")
	signal.InitSignal()
}
