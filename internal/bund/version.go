package bund

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/banner"
	"github.com/vulogov/Bund/internal/conf"
)

func Version() {
	Init()
	log.Debug("[ BUND ] bund.Version() is reached")
	banner.Banner(fmt.Sprintf("[ BUND %v ]", conf.EVersion))
	banner.Table()
}
