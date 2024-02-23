// http_logging.go
package jamfpro

import "log"

// Logger is an interface for logging within the SDK.// LogLevel mirrors the httpclient's LogLevel to maintain compatibility.
type LogLevel int

// Exporting LogLevel constants matching httpclient package.
const (
	LogLevelNone LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

// Logger interface to match httpclient's Logger interface.
type Logger interface {
	SetLevel(level LogLevel)
	Trace(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

// defaultLogger now includes logLevel to control the logging output.
type defaultLogger struct {
	logLevel LogLevel
}

// SetLevel sets the log level for the logger.
func (d *defaultLogger) SetLevel(level LogLevel) {
	d.logLevel = level
}

// Trace logs if the level is set to LogLevelDebug or lower.
func (d *defaultLogger) Trace(msg string, keysAndValues ...interface{}) {
	if d.logLevel >= LogLevelDebug {
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

// NewDefaultLogger returns a new default logger instance with a default log level.
func NewDefaultLogger() Logger {
	return &defaultLogger{
		logLevel: LogLevelWarning, // Default level can be set here.
	}
}

// Expose the default logger instance for use in the main package.
var DefaultLogger = NewDefaultLogger()
