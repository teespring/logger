package logger

import (
	"io"
	"os"
	"strings"
	"time"
)

var (
	out                 io.Writer        = os.Stderr
	verbosity           int              = Verbosity()
	enabled, allEnabled                  = Enabled()
	now                 func() time.Time = time.Now
)

func Enabled() (map[string]bool, bool) {
	val := os.Getenv("LOG")

	if val == "*" {
		return nil, true
	}

	all := map[string]bool{}
	keys := strings.Split(val, ",")

	for _, key := range keys {
		all[key] = true
	}

	return all, false
}

func IsEnabled(name string) bool {
	if allEnabled {
		return true
	}

	_, ok := enabled[name]
	return ok
}

func Verbosity() int {
	level := os.Getenv("LOG_LEVEL")

	if strings.ToUpper(level) == "TIMER" {
		return 2
	}

	if strings.ToUpper(level) == "ERROR" {
		return 3
	}

	return 1
}

func SetOutput(w *os.File) {
	colorEnabled = false
	out = w
}
