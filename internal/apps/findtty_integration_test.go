package apps_test

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/gnomos-unix/findtty/internal/apps"
	"github.com/gnomos-unix/findtty/internal/bufferio"
	"golang.org/x/exp/slices"
)

const (
	dataFolder = "../../test/data"
	tmpFolder  = "../../test/tmp"
)

func TestFindTTYApp(t *testing.T) {
	type input struct {
		path           string
		beforeScanFunc func(t *testing.T)
	}
	type output struct {
		ttyFiles []string
	}
	type testCase struct {
		name string
		in   input
		out  output
	}

	tests := []testCase{
		{
			name: "should return new terminal device when found terminal device",
			in: input{
				path: dataFolder,
				beforeScanFunc: func(t *testing.T) {
					moveTempTTYFile(t, tmpFolder, dataFolder)
				},
			},
			out: output{
				ttyFiles: []string{
					"../../test/data/tmpTTY",
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			in := fakeInput(t)
			defer in.Close()

			scanner := bufferio.Scanner{
				Input: bufio.NewScanner(in),
				BeforeScanFunc: func() {
					tc.in.beforeScanFunc(t)
				},
			}
			app := apps.NewFindTTY(tc.in.path, false, scanner)
			app.Run()
			moveTempTTYFile(t, dataFolder, tmpFolder)

			if len(tc.out.ttyFiles) != len(app.NewTTYFiles) {
				t.Errorf("unexpected quatity of new terminal devices: wants %d, got %d", len(tc.out.ttyFiles), len(app.NewTTYFiles))
				t.FailNow()
			}

			for _, file := range app.NewTTYFiles {
				if !slices.Contains(tc.out.ttyFiles, file) {
					t.Errorf("terminal device not found: %s", file)
					t.FailNow()
				}
			}
		})
	}
}

func fakeInput(t *testing.T) *os.File {
	in, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.WriteString(in, "\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, io.SeekCurrent|io.SeekEnd|io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}

	return in
}

func moveTempTTYFile(t *testing.T, src, dest string) {
	err := os.Rename(src+"/tmpTTY", dest+"/tmpTTY")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
		return
	}
}
