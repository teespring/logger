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

	switch strings.ToUpper(level) {
	case "INFO":
		return VERBOSITY_INFO
	case "TIMER":
		return VERBOSITY_TIMER
	case "WARN":
		return VERBOSITY_WARN
	case "ERROR":
		return VERBOSITY_ERROR
	}

	return VERBOSITY_DEBUG
}

func SetOutput(w io.Writer) {
	colorEnabled = false
	out = w
}
