package apps

import (
	"io/fs"
	"path/filepath"
	"syscall"

	"github.com/gnomos-unix/findtty/internal/bufferio"
	"github.com/gnomos-unix/findtty/internal/logs"
	"golang.org/x/term"
)

// FindTTY is a representation of findtty application
type FindTTY struct {
	NewTTYFiles []string
	Scanner     bufferio.Scanner
	RootPath    string
	Logger      logs.Logger

	ttyFiles map[string]struct{}
}

// PrintTTYFiles shows all terminal devices found
func (app FindTTY) PrintTTYFiles() {
	if len(app.NewTTYFiles) == 0 {
		app.Logger.Info("there isn't any new terminal device available.\n")
		return
	}
	app.Logger.Info("new terminal device(s) found:\n")
	for _, f := range app.NewTTYFiles {
		app.Logger.Info("- %s\n", f)
	}
}

// Run runs the findtty application in order to identify new terminal devices connected in your computer
func (app *FindTTY) Run() {
	app.walk(func(path string) error {
		app.ttyFiles[path] = struct{}{}
		app.Logger.Debug("terminal device '%s' loaded\n", path)
		return nil
	})

	app.Logger.Info("Connect your device and press enter.\n")
	app.Scanner.Scan()
	// app.Scanner.Scan()
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()

	app.walk(func(path string) error {
		_, ok := app.ttyFiles[path]
		if !ok {
			app.Logger.Debug("new terminal device found: '%s'\n", path)
			app.NewTTYFiles = append(app.NewTTYFiles, path)
		}
		return nil
	})
}

func (app FindTTY) walk(findTTYFunc func(path string) error) {
	filepath.Walk(app.RootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fd := app.openFile(path, info.IsDir())
		if fd == -1 {
			return nil
		}

		if term.IsTerminal(fd) {
			return findTTYFunc(path)
		}
		return nil
	})
}

func (app FindTTY) openFile(path string, isDir bool) int {
	if isDir {
		app.Logger.Error("path '%s' is not a valid file\n", path)
		return -1
	}

	fd, err := syscall.Open(path, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0600)
	if err != nil {
		app.Logger.Error("it is not possible to open file '%s': %s\n", path, err.Error())
		return -1
	}

	return fd
}
