package bund

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/vulogov/Bund/internal/conf"
)

func Main() {
	switch kingpin.MustParse(conf.App.Parse(os.Args[1:])) {
	case conf.Version.FullCommand():
		Version()
	case conf.Shell.FullCommand():
		Shell()
	case conf.Run.FullCommand():
		Run()
	case conf.Eval.FullCommand():
		Eval()

	}
	Fin()
}
