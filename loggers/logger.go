package loggers

import (
	"log"
	"os"
)

// Logger is the app logger
type Logger interface {
	Info(message string)
	Warn(message string)
	Error(message string)
}

type logger struct {
	file string
	Il   *log.Logger
	Wl   *log.Logger
	El   *log.Logger
}

// GetLogger give one logger
func GetLogger(filename string) Logger {
	return &logger{
		file: filename,
		Il:   log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		Wl:   log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		El:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
	}
}

// Info log an info
func (l *logger) Info(message string) {
	l.Il.Println(l.file, message)
}

// Warn log a Warn
func (l *logger) Warn(message string) {
	l.Wl.Println(l.file, message)
}

// Error log an error
func (l *logger) Error(message string) {
	l.El.Println(l.file, message)
}
