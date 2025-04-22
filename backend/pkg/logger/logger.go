package logger

import (
	"log"
	"os"
)

// Logger is a simple logging interface
type Logger interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

// SimpleLogger is a basic implementation of the Logger interface
type SimpleLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

// NewLogger creates a new SimpleLogger
func NewLogger() Logger {
	return &SimpleLogger{
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Debug logs a debug message
func (l *SimpleLogger) Debug(message string) {
	l.debugLogger.Println(message)
}

// Info logs an info message
func (l *SimpleLogger) Info(message string) {
	l.infoLogger.Println(message)
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(message string) {
	l.warnLogger.Println(message)
}

// Error logs an error message
func (l *SimpleLogger) Error(message string) {
	l.errorLogger.Println(message)
}

// Fatal logs a fatal message and exits the application
func (l *SimpleLogger) Fatal(message string) {
	l.fatalLogger.Fatalln(message)
}
