package bufferio

import (
	"bufio"
	"os"
)

type Scanner struct {
	Input          *bufio.Scanner
	BeforeScanFunc func()
}

func NewScanner() Scanner {
	return Scanner{
		Input:          bufio.NewScanner(os.Stdin),
		BeforeScanFunc: func() {},
	}
}

func (s Scanner) Scan() {
	s.BeforeScanFunc()
	s.Input.Scan()
}
