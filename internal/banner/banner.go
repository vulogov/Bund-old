package banner

import (
	"github.com/common-nighthawk/go-figure"

	"github.com/vulogov/Bund/internal/conf"
)

func Banner(txt string) {
	if *conf.VBanner {
		if *conf.Color {
			b := figure.NewColorFigure(txt, "", "gray", false)
			b.Print()
		} else {
			b := figure.NewFigure(txt, "", true)
			b.Print()
		}
	}
}
