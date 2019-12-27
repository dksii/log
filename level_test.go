package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	ttable := []struct {
		level         string
		expectedLevel Level
		expectedErr   error
	}{
		{level: "error", expectedLevel: ErrorLevel, expectedErr: nil},
		{level: "warn", expectedLevel: WarnLevel, expectedErr: nil},
		{level: "warning", expectedLevel: WarnLevel, expectedErr: nil},
		{level: "info", expectedLevel: InfoLevel, expectedErr: nil},
		{level: "debug", expectedLevel: DebugLevel, expectedErr: nil},
		{level: "unknown", expectedLevel: 0, expectedErr: errors.New("unknown level")},
	}

	for _, tt := range ttable {
		tt := tt

		t.Run(tt.level, func(t *testing.T) {
			if tt.expectedErr != nil {
				_, actualErr := ParseLevel(tt.level)
				assert.Equal(t, tt.expectedErr, actualErr)
				return
			}

			actualLevel, err := ParseLevel(tt.level)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedLevel, actualLevel)
		})
	}
}

func TestLevel_correct(t *testing.T) {
	assert.True(t, ErrorLevel.correct())
	assert.True(t, WarnLevel.correct())
	assert.True(t, InfoLevel.correct())
	assert.True(t, DebugLevel.correct())
	assert.False(t, Level(4).correct())
	assert.False(t, Level(100).correct())
}

func TestLevel_bytes(t *testing.T) {
	assert.Equal(t, []byte("ERR"), ErrorLevel.bytes())
	assert.Equal(t, []byte("WRN"), WarnLevel.bytes())
	assert.Equal(t, []byte("INF"), InfoLevel.bytes())
	assert.Equal(t, []byte("DBG"), DebugLevel.bytes())
}
