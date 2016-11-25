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
	Il *log.Logger
	Wl *log.Logger
	El *log.Logger
}

// New init one logger
func New() Logger {
	return &logger{
		Il: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		Wl: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		El: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
	}
}

// Info log an info
func (l *logger) Info(message string) {
	l.Il.Println(message)
}

// Warn log a Warn
func (l *logger) Warn(message string) {
	l.Wl.Println(message)
}

// Error log an error
func (l *logger) Error(message string) {
	l.El.Println(message)
}
