package logs

import (
	"fmt"
	"os"
)

type Logger struct {
	IsVerbose bool
}

// Info print in stdout always
func (log Logger) Info(format string, args ...any) {
	fmt.Fprintf(os.Stdout, format, args...)
}

// Debug print in stdout just when the verbose is set to true
func (log Logger) Debug(format string, args ...any) {
	if log.IsVerbose {
		fmt.Fprintf(os.Stdout, format, args...)
	}
}

// Error print in stderr just when the verbose is set to true
func (log Logger) Error(format string, args ...any) {
	if log.IsVerbose {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}
