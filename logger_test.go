package logger

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	var expected = "This is a test message.\n"

	// Overwrite settings
	var buff = &bytes.Buffer{}
	SetOutput(buff)
	allEnabled = true

	// Create a logger to test
	var subject = New("Test Logger")

	subject.Write("This is a test message.")

	if s := buff.String(); !strings.EqualFold(s, expected) {
		t.Errorf("The expected string did not get written: Expected: \"%s\"; Received: \"%s\"", expected, s)
	}
}

func TestOutput(t *testing.T) {
	var expected = "{ \"time\":\"2016-06-13 11:59:50.000123456 -0700 PDT\", \"package\":\"Test Logger\"," +
		" \"level\":\"VERB\", \"msg\":\"This is a test message.\" }\n"

	// Overwrite settings
	var buff = &bytes.Buffer{}
	SetOutput(buff)
	allEnabled = true
	now = func() time.Time {
		return time.Unix(1465844390, 123456)
	}

	// Create a logger to test
	var subject = New("Test Logger")
	var attrs = &Attrs{}

	subject.Output(33, "VERB", "This is a test message.", attrs)

	if s := buff.String(); !strings.EqualFold(s, expected) {
		t.Errorf("The expected string did not get written: Expected: \"%s\"; Received: \"%s\"", expected, s)
	}
}

func TestDebug(t *testing.T) {
	var expected = "{ \"time\":\"2016-06-13 11:59:50.000123456 -0700 PDT\", \"package\":\"Test Logger\"," +
		" \"level\":\"DEBUG\", \"msg\":\"This is a test message.\" }\n"

	// Overwrite output writer w/ Buffer
	var buff = &bytes.Buffer{}
	SetOutput(buff)
	allEnabled = true
	now = func() time.Time {
		return time.Unix(1465844390, 123456)
	}

	// Create a logger to test
	var subject = New("Test Logger")

	subject.Debug("This is a test message.")

	if s := buff.String(); !strings.EqualFold(s, expected) {
		t.Errorf("The expected string did not get written: Expected: \"%s\"; Received: \"%s\"", expected, s)
	}
}

func TestInfo(t *testing.T) {
	var expected = "{ \"time\":\"2016-06-13 11:59:50.000123456 -0700 PDT\", \"package\":\"Test Logger\"," +
		" \"level\":\"INFO\", \"msg\":\"This is a test message.\" }\n"

	// Overwrite output writer w/ Buffer
	var buff = &bytes.Buffer{}
	SetOutput(buff)
	allEnabled = true
	now = func() time.Time {
		return time.Unix(1465844390, 123456)
	}

	// Create a logger to test
	var subject = New("Test Logger")

	subject.Info("This is a test message.")

	if s := buff.String(); !strings.EqualFold(s, expected) {
		t.Errorf("The expected string did not get written: Expected: \"%s\"; Received: \"%s\"", expected, s)
	}
}

func TestWarn(t *testing.T) {
	var expected = "{ \"time\":\"2016-06-13 11:59:50.000123456 -0700 PDT\", \"package\":\"Test Logger\"," +
		" \"level\":\"WARN\", \"msg\":\"This is a test message.\" }\n"

	// Overwrite output writer w/ Buffer
	var buff = &bytes.Buffer{}
	SetOutput(buff)
	allEnabled = true
	now = func() time.Time {
		return time.Unix(1465844390, 123456)
	}

	// Create a logger to test
	var subject = New("Test Logger")

	subject.Warn("This is a test message.")

	if s := buff.String(); !strings.EqualFold(s, expected) {
		t.Errorf("The expected string did not get written: Expected: \"%s\"; Received: \"%s\"", expected, s)
	}
}

func TestError(t *testing.T) {
	var expected = "{ \"time\":\"2016-06-13 11:59:50.000123456 -0700 PDT\", \"package\":\"Test Logger\"," +
		" \"level\":\"ERROR\", \"msg\":\"This is a test message.\" }\n"

	// Overwrite output writer w/ Buffer
	var buff = &bytes.Buffer{}
	SetOutput(buff)
	allEnabled = true
	now = func() time.Time {
		return time.Unix(1465844390, 123456)
	}

	// Create a logger to test
	var subject = New("Test Logger")

	subject.Error("This is a test message.")

	if s := buff.String(); !strings.EqualFold(s, expected) {
		t.Errorf("The expected string did not get written: Expected: \"%s\"; Received: \"%s\"", expected, s)
	}
}
