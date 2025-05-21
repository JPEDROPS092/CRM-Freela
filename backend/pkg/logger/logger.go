package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Logger is a simple logging interface
type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
	RequestInfo(method, path, ip, status string, latency time.Duration)
	APICall(endpoint, method, status string, latency time.Duration)
}

// SimpleLogger is a basic implementation of the Logger interface
type SimpleLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	apiLogger   *log.Logger
	reqLogger   *log.Logger
}

// LogLevel represents the logging level
type LogLevel string

const (
	// DebugLevel logs everything
	DebugLevel LogLevel = "debug"
	// InfoLevel logs info, warnings, errors, and fatal
	InfoLevel LogLevel = "info"
	// WarnLevel logs warnings, errors, and fatal
	WarnLevel LogLevel = "warn"
	// ErrorLevel logs errors and fatal
	ErrorLevel LogLevel = "error"
	// FatalLevel logs only fatal
	FatalLevel LogLevel = "fatal"
)

// NewLogger creates a new SimpleLogger
func NewLogger() Logger {
	return &SimpleLogger{
		debugLogger: log.New(os.Stdout, "\033[36mDEBUG\033[0m: ", log.Ldate|log.Ltime),
		infoLogger:  log.New(os.Stdout, "\033[32mINFO\033[0m:  ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "\033[33mWARN\033[0m:  ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "\033[31mERROR\033[0m: ", log.Ldate|log.Ltime),
		fatalLogger: log.New(os.Stderr, "\033[35mFATAL\033[0m: ", log.Ldate|log.Ltime),
		apiLogger:   log.New(os.Stdout, "\033[34mAPI\033[0m:   ", log.Ldate|log.Ltime),
		reqLogger:   log.New(os.Stdout, "\033[32mREQ\033[0m:   ", log.Ldate|log.Ltime),
	}
}

// formatMessage formats a message with optional arguments
func formatMessage(message string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(message, args...)
	}
	return message
}

// getCallerInfo returns the file and line number of the caller
func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	// Get only the filename without the full path
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return fmt.Sprintf("[%s:%d] ", short, line)
}

// Debug logs a debug message
func (l *SimpleLogger) Debug(message string, args ...interface{}) {
	callerInfo := getCallerInfo()
	l.debugLogger.Println(callerInfo + formatMessage(message, args...))
}

// Info logs an info message
func (l *SimpleLogger) Info(message string, args ...interface{}) {
	l.infoLogger.Println(formatMessage(message, args...))
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(message string, args ...interface{}) {
	l.warnLogger.Println(formatMessage(message, args...))
}

// Error logs an error message
func (l *SimpleLogger) Error(message string, args ...interface{}) {
	callerInfo := getCallerInfo()
	l.errorLogger.Println(callerInfo + formatMessage(message, args...))
}

// Fatal logs a fatal message and exits the application
func (l *SimpleLogger) Fatal(message string, args ...interface{}) {
	callerInfo := getCallerInfo()
	l.fatalLogger.Println(callerInfo + formatMessage(message, args...))
	os.Exit(1)
}

// RequestInfo logs information about an HTTP request
func (l *SimpleLogger) RequestInfo(method, path, ip, status string, latency time.Duration) {
	l.reqLogger.Printf("%s %s | %s | %s | %s", 
		method, 
		path, 
		ip, 
		status,
		latency.String())
}

// APICall logs information about an API call
func (l *SimpleLogger) APICall(endpoint, method, status string, latency time.Duration) {
	l.apiLogger.Printf("%s %s | Status: %s | Latency: %s", 
		method, 
		endpoint, 
		status,
		latency.String())
}
