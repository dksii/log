package log

import (
	"io"
	"os"
)

// standard logger
var std, _ = New(os.Stderr, InfoLevel)

// SetOutput sets output of the standard logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetLevel sets log level of the standard logger.
func SetLevel(new Level) error {
	return std.SetLevel(new)
}

// GetLevel returns current log level of the standard logger.
func GetLevel() Level {
	return std.GetLevel()
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
