package logger

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func New() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile),
		warn:  log.New(os.Stdout, "[WARN]  ", log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.info.Printf(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.warn.Printf(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.error.Printf(msg, args...)
}
