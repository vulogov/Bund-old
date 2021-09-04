package banner

import (
	"os"

	"github.com/mgutz/ansi"
	"github.com/tomlazar/table"

	"github.com/vulogov/Bund/internal/conf"
)

func Table() {
	var cfg table.Config

	if !*conf.VTable {
		return
	}

	cfg.ShowIndex = true
	if *conf.Color {
		cfg.Color = true
		cfg.AlternateColors = true
		cfg.TitleColorCode = ansi.ColorCode("white+buf")
		cfg.AltColorCodes = []string{"", ansi.ColorCode("white:grey+h")}
	} else {
		cfg.Color = false
		cfg.AlternateColors = false
		cfg.TitleColorCode = ansi.ColorCode("white+buf")
		cfg.AltColorCodes = []string{"", ansi.ColorCode("white:grey+h")}
	}
	if *conf.VTable {
		tab := table.Table{
			Headers: []string{"Description", "Value"},
			Rows: [][]string{
				{"Version", conf.EVersion},
				{"Extended version", conf.BVersion},
				{"Args", *conf.Args},
			},
		}
		tab.WriteTable(os.Stdout, &cfg)
	}
}
