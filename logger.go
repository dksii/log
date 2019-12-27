package log

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

// Logger represents an active logging object.
type Logger struct {
	out   io.Writer
	level Level
	mu    sync.Mutex
	buf   []byte
}

// New creates Logger instance.
func New(out io.Writer, level Level) (*Logger, error) {
	if !level.correct() {
		return nil, errors.New("incorrect level")
	}

	return &Logger{out: out, level: level}, nil
}

// SetOutput sets the output destination.
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	l.out = w
	l.mu.Unlock()
}

// SetLevel sets current log level.
func (l *Logger) SetLevel(new Level) error {
	if !new.correct() {
		return errors.New("incorrect level")
	}

	atomic.StoreUint32((*uint32)(&l.level), uint32(new))
	return nil
}

// GetLevel returns current log level.
func (l *Logger) GetLevel() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(args ...interface{}) {
	l.log(DebugLevel, args...)
}

// Debugf logs a message at level Debug.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(DebugLevel, format, args...)
}

// Info logs a message at level Info.
func (l *Logger) Info(args ...interface{}) {
	l.log(InfoLevel, args...)
}

// Infof logs a message at level Info.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(InfoLevel, format, args...)
}

// Warn logs a message at level Warn.
func (l *Logger) Warn(args ...interface{}) {
	l.log(WarnLevel, args...)
}

// Warnf logs a message at level Warn.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logf(WarnLevel, format, args...)
}

// Error logs a message at level Error.
func (l *Logger) Error(args ...interface{}) {
	l.log(ErrorLevel, args...)
}

// Errorf logs a message at level Error.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(ErrorLevel, format, args...)
}

func (l *Logger) supportsLevel(level Level) bool {
	return l.GetLevel() >= level
}

func (l *Logger) log(level Level, args ...interface{}) {
	// exit early to not use fmt.Sprint if level not supported
	if !l.supportsLevel(level) {
		return
	}
	l.output(level, fmt.Sprint(args...))
}

func (l *Logger) logf(level Level, format string, args ...interface{}) {
	// exit early to not use fmt.Sprintf if level not supported
	if !l.supportsLevel(level) {
		return
	}
	l.output(level, fmt.Sprintf(format, args...))
}

func (l *Logger) output(level Level, s string) {
	// save current time for logging
	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()

	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, level)
	l.buf = append(l.buf, s...)

	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, _ = l.out.Write(l.buf)
}

// we assume this func protected by mutex
func (l *Logger) formatHeader(buf *[]byte, t time.Time, level Level) {
	// append date part
	year, month, day := t.Date()
	itoa(buf, year, 4)
	*buf = append(*buf, '/')
	itoa(buf, int(month), 2)
	*buf = append(*buf, '/')
	itoa(buf, day, 2)
	*buf = append(*buf, ' ')
	// append hours, minutes, seconds
	hour, min, sec := t.Clock()
	itoa(buf, hour, 2)
	*buf = append(*buf, ':')
	itoa(buf, min, 2)
	*buf = append(*buf, ':')
	itoa(buf, sec, 2)
	// append milliseconds
	*buf = append(*buf, '.')
	itoa(buf, t.Nanosecond()/1e6, 3)
	// append log level
	*buf = append(*buf, ' ', '[')
	*buf = append(*buf, level.bytes()...)
	*buf = append(*buf, ']', ' ')
}
