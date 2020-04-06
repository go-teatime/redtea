package log

import (
	"io"
	"log"
	"os"
	"sync/atomic"
)

type Level uint32

const (
	LevelNone Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelBuss
)

type Logger struct {
	level  Level
	logger *log.Logger
}

func NewLogger() *Logger {
	out := os.Stderr
	return &Logger{
		level:  LevelDebug,
		logger: log.New(out, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.level), uint32(level))
}

func (l *Logger) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

func (l *Logger) Trace(format string, a ...interface{}) {
	if l.Level() <= LevelTrace {
		l.logger.Printf("[TRACE] "+format, a...)
	}
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if l.Level() <= LevelDebug {
		l.logger.Printf("[DEBUG] "+format, a...)
	}
}

func (l *Logger) Info(format string, a ...interface{}) {
	if l.Level() <= LevelInfo {
		l.logger.Printf("[INFO] "+format, a...)
	}
}

func (l *Logger) Warn(format string, a ...interface{}) {
	if l.Level() <= LevelWarn {
		l.logger.Printf("[WARN] "+format, a...)

	}
}
func (l *Logger) Error(format string, a ...interface{}) {
	if l.Level() <= LevelError {
		l.logger.Printf("[ERROR] "+format, a...)
	}
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	if l.Level() <= LevelFatal {
		l.logger.Printf("[FATAL] "+format, a...)
	}
}

func (l *Logger) Buss(format string, a ...interface{}) {
	if l.Level() <= LevelBuss {
		l.logger.Printf("[BUSS] "+format, a...)
	}
}
