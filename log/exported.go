package log

import (
	"io"
	"os"
	"strings"
)

var std = NewLogger()

func init() {
	fOut := NewLogByDayWriter("./logfiles", "")
	if fOut != nil {
		std.SetOutput(io.MultiWriter(os.Stderr, fOut))
	}
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetLevel(lstr string) {
	m := map[string]Level{
		"trace": LevelTrace,
		"debug": LevelDebug,
		"info":  LevelInfo,
		"warn":  LevelWarn,
		"error": LevelError,
		"fatal": LevelFatal,
		"buss":  LevelBuss,
	}
	s := strings.ToLower(lstr)
	l, ok := m[s]
	if ok {
		std.SetLevel(l)
	}
}

func Trace(format string, a ...interface{}) {
	std.Trace(format, a...)
}

func Info(format string, a ...interface{}) {
	std.Info(format, a...)
}

func Debug(format string, a ...interface{}) {
	std.Debug(format, a...)
}

func Warn(format string, a ...interface{}) {
	std.Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	std.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	std.Fatal(format, a...)
}

func Buss(format string, a ...interface{}) {
	std.Buss(format, a...)
}
