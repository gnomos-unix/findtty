package apps

import (
	"github.com/gnomos-unix/findtty/internal/bufferio"
	"github.com/gnomos-unix/findtty/internal/logs"
)

func NewFindTTY(path string, verbose bool, scanner bufferio.Scanner) FindTTY {
	app := FindTTY{
		NewTTYFiles: []string{},
		Scanner:     scanner,
		RootPath:    path,

		ttyFiles: map[string]struct{}{},
		Logger: logs.Logger{
			IsVerbose: verbose,
		},
	}
	return app
}
