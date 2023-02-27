package main

import (
	"flag"

	"github.com/gnomos-unix/findtty/internal/apps"
	"github.com/gnomos-unix/findtty/internal/bufferio"
)

func main() {
	var path string
	var verbose bool

	flag.StringVar(&path, "path", "/dev", "set the root path where will be used to search new terminal files")
	flag.BoolVar(&verbose, "verbose", false, "set verbose (default false)")
	flag.Parse()

	app := apps.NewFindTTY(path, verbose, bufferio.NewScanner())

	app.Run()
	app.PrintTTYFiles()
}
