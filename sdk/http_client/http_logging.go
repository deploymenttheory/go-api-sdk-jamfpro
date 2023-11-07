// http_logging.go
package http_client

import "log"

type LogLevel int

const (
	LogLevelNone LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

// Logger is an interface for logging within the SDK.
type Logger interface {
	SetLevel(level LogLevel)
	Trace(msg string, keysAndValues ...interface{}) // For very detailed logs
	Debug(msg string, keysAndValues ...interface{}) // For development and troubleshooting
	Info(msg string, keysAndValues ...interface{})  // Informational messages
	Warn(msg string, keysAndValues ...interface{})  // For potentially problematic situations
	Error(msg string, keysAndValues ...interface{}) // For errors that might still allow the app to continue running
	Fatal(msg string, keysAndValues ...interface{}) // For errors that might prevent the app from continuing
}

// defaultLogger is the default logger based on Go's standard log package and includes a logLevel field to keep track of the current logging level.
type defaultLogger struct {
	logLevel LogLevel
}

// SetLevel sets the current logging level for the defaultLogger.
func (d *defaultLogger) SetLevel(level LogLevel) {
	d.logLevel = level
}

// NewDefaultLogger now initializes a defaultLogger with a default log level.
func NewDefaultLogger() Logger {
	return &defaultLogger{
		logLevel: LogLevelWarning, // default log level.
	}
}

// Trace checks if the current log level permits debug messages before logging.
func (d *defaultLogger) Trace(msg string, keysAndValues ...interface{}) {
	if d.logLevel >= LogLevelDebug { // Trace is a part of LogLevelDebug
		log.Println("[TRACE]", msg, keysAndValues)
	}
}

// Debug checks if the current log level permits debug messages before logging.
func (d *defaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	if d.logLevel >= LogLevelDebug {
		log.Println("[DEBUG]", msg, keysAndValues)
	}
}

// Info checks if the current log level permits debug messages before logging.
func (d *defaultLogger) Info(msg string, keysAndValues ...interface{}) {
	if d.logLevel >= LogLevelInfo {
		log.Println("[INFO]", msg, keysAndValues)
	}
}

// Warn checks if the current log level permits Warning messages before logging.
func (d *defaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	if d.logLevel >= LogLevelWarning {
		log.Println("[WARN]", msg, keysAndValues)
	}
}

// Error checks if the current log level is greater than LogLevelNone, before logging.
func (d *defaultLogger) Error(msg string, keysAndValues ...interface{}) {
	if d.logLevel > LogLevelNone {
		log.Println("[ERROR]", msg, keysAndValues)
	}
}

// Fatal checks if the current log level is greater than LogLevelNone, before logging.
func (d *defaultLogger) Fatal(msg string, keysAndValues ...interface{}) {
	if d.logLevel > LogLevelNone {
		log.Fatalln("[FATAL]", msg, keysAndValues)
	}
}
