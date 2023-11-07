// http_logging.go
package http_client

import "log"

// Logger is an interface for logging within the SDK.
type Logger interface {
	Trace(msg string, keysAndValues ...interface{}) // For very detailed logs
	Debug(msg string, keysAndValues ...interface{}) // For development and troubleshooting
	Info(msg string, keysAndValues ...interface{})  // Informational messages
	Warn(msg string, keysAndValues ...interface{})  // For potentially problematic situations
	Error(msg string, keysAndValues ...interface{}) // For errors that might still allow the app to continue running
	Fatal(msg string, keysAndValues ...interface{}) // For errors that might prevent the app from continuing
}

// defaultLogger is the default logger based on Go's standard log package.
type defaultLogger struct{}

func (d *defaultLogger) Trace(msg string, keysAndValues ...interface{}) {
	log.Println("[TRACE]", msg, keysAndValues)
}

func (d *defaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.Println("[DEBUG]", msg, keysAndValues)
}

func (d *defaultLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Println("[INFO]", msg, keysAndValues)
}

func (d *defaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	log.Println("[WARN]", msg, keysAndValues)
}

func (d *defaultLogger) Error(msg string, keysAndValues ...interface{}) {
	log.Println("[ERROR]", msg, keysAndValues)
}

func (d *defaultLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatalln("[FATAL]", msg, keysAndValues)
}

// NewDefaultLogger returns a new instance of the default logger.
func NewDefaultLogger() Logger {
	return &defaultLogger{}
}
