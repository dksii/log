package log

import "errors"

// Level type.
type Level uint32

// Available levels.
const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

// ParseLevel parses log level.
func ParseLevel(level string) (Level, error) {
	switch level {
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	default:
		return 0, errors.New("unknown level")
	}
}

// bytes representation
var levelBytes = [][]byte{
	[]byte("ERR"),
	[]byte("WRN"),
	[]byte("INF"),
	[]byte("DBG"),
}

func (l Level) bytes() []byte {
	return levelBytes[l]
}

func (l Level) correct() bool {
	return l <= DebugLevel
}
