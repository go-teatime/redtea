package log

import "testing"

func TestLog(t *testing.T) {
	Trace("hello%s", "world")
	Debug("hello%s", "world")
	Info("hello%s", "world")
	Warn("hello%s", "world")
	Error("hello%s", "world")
	Fatal("hello%s", "world")
	Buss("hello%s", "world")
}
