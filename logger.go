package logger

import (
	"fmt"
)

const (
	LEVEL_DEBUG     = "DEBUG"
	LEVEL_INFO      = "INFO"
	LEVEL_TIMER     = "TIMER"
	LEVEL_WARN      = "WARN"
	LEVEL_ERROR     = "ERROR"
	VERBOSITY_DEBUG = 1
	VERBOSITY_TIMER = 2
	VERBOSITY_INFO  = 3
	VERBOSITY_WARN  = 4
	VERBOSITY_ERROR = 5
)

type Logger struct {
	Name      string
	IsEnabled bool
	Color     string
}

func New(name string) *Logger {
	return &Logger{
		Name:      name,
		IsEnabled: IsEnabled(name),
		Color:     nextColor(),
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if verbosity > VERBOSITY_DEBUG {
		return
	}

	if !l.IsEnabled {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(VERBOSITY_DEBUG, LEVEL_DEBUG, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Info(format string, v ...interface{}) {
	if verbosity > VERBOSITY_INFO {
		return
	}

	if !l.IsEnabled {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(VERBOSITY_INFO, LEVEL_INFO, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Timer() *Timer {
	return &Timer{
		Logger:    l,
		Start:     Now(),
		IsEnabled: l.IsEnabled && verbosity < 3,
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if verbosity > VERBOSITY_WARN {
		return
	}

	if !l.IsEnabled {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(VERBOSITY_WARN, LEVEL_WARN, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Error(format string, v ...interface{}) {
	if !l.IsEnabled {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(VERBOSITY_ERROR, LEVEL_ERROR, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Output(verbosity int, sort string, msg string, attrs *Attrs) {
	l.Write(l.Format(verbosity, sort, msg, attrs))
}

func (l *Logger) Write(log string) {
	fmt.Fprintln(out, log)
}
