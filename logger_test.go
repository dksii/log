package log

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		logger, err := New(ioutil.Discard, ErrorLevel)
		assert.NoError(t, err)
		assert.NotNil(t, logger)
	})

	t.Run("fail", func(t *testing.T) {
		_, err := New(ioutil.Discard, Level(4))
		assert.Error(t, err)
		assert.EqualError(t, err, "incorrect level")
	})
}

func TestLogger_SetOutput(t *testing.T) {
	out1, out2 := os.Stdout, os.Stderr
	logger, err := New(ioutil.Discard, ErrorLevel)

	assert.NoError(t, err)
	assert.NotEqual(t, logger.out, out1)
	assert.NotEqual(t, logger.out, out2)

	logger.SetOutput(out1)
	assert.Equal(t, logger.out, out1)
	assert.NotEqual(t, logger.out, out2)

	logger.SetOutput(out2)
	assert.NotEqual(t, logger.out, out1)
	assert.Equal(t, logger.out, out2)
}

func TestLogger_GetLevel(t *testing.T) {
	logger, err := New(ioutil.Discard, ErrorLevel)
	assert.NoError(t, err)
	assert.Equal(t, ErrorLevel, logger.GetLevel())
}

func TestLogger_SetLevel(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		logger, err := New(ioutil.Discard, ErrorLevel)
		assert.NoError(t, err)

		assert.Equal(t, ErrorLevel, logger.GetLevel())
		err = logger.SetLevel(InfoLevel)
		assert.NoError(t, err)
		assert.Equal(t, InfoLevel, logger.GetLevel())
	})

	t.Run("fail", func(t *testing.T) {
		logger, err := New(ioutil.Discard, ErrorLevel)
		assert.NoError(t, err)
		err = logger.SetLevel(Level(4))
		assert.EqualError(t, err, "incorrect level")
	})
}

func TestLogger_Debug(t *testing.T) {
	t.Run("skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, InfoLevel)
		assert.NoError(t, err)
		logger.Debug("message")
		assert.Zero(t, buf.Len())
	})

	t.Run("not skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, DebugLevel)
		assert.NoError(t, err)
		logger.Debug("message")
		assert.True(t, strings.HasSuffix(buf.String(), "[DBG] message\n"))
	})
}

func TestLogger_Debugf(t *testing.T) {
	t.Run("skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, InfoLevel)
		assert.NoError(t, err)
		logger.Debugf("message with %s", "arg")
		assert.Zero(t, buf.Len())
	})

	t.Run("not skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, DebugLevel)
		assert.NoError(t, err)
		logger.Debugf("message with %s", "arg")
		assert.True(t, strings.HasSuffix(buf.String(), "[DBG] message with arg\n"))
	})
}

func TestLogger_Info(t *testing.T) {
	t.Run("skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, WarnLevel)
		assert.NoError(t, err)
		logger.Info("message")
		assert.Zero(t, buf.Len())
	})

	t.Run("not skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, InfoLevel)
		assert.NoError(t, err)
		logger.Info("message")
		assert.True(t, strings.HasSuffix(buf.String(), "[INF] message\n"))
	})
}

func TestLogger_Infof(t *testing.T) {
	t.Run("skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, ErrorLevel)
		assert.NoError(t, err)
		logger.Infof("message with %s", "arg")
		assert.Zero(t, buf.Len())
	})

	t.Run("not skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, DebugLevel)
		assert.NoError(t, err)
		logger.Infof("message with %s", "arg")
		assert.True(t, strings.HasSuffix(buf.String(), "[INF] message with arg\n"))
	})
}

func TestLogger_Warn(t *testing.T) {
	t.Run("skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, ErrorLevel)
		assert.NoError(t, err)
		logger.Warn("message")
		assert.Zero(t, buf.Len())
	})

	t.Run("not skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, WarnLevel)
		assert.NoError(t, err)
		logger.Warn("message")
		assert.True(t, strings.HasSuffix(buf.String(), "[WRN] message\n"))
	})
}

func TestLogger_Warnf(t *testing.T) {
	t.Run("skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, ErrorLevel)
		assert.NoError(t, err)
		logger.Warnf("message with %s", "arg")
		assert.Zero(t, buf.Len())
	})

	t.Run("not skipped", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger, err := New(buf, InfoLevel)
		assert.NoError(t, err)
		logger.Warnf("message with %s", "arg")
		assert.True(t, strings.HasSuffix(buf.String(), "[WRN] message with arg\n"))
	})
}

func TestLogger_Error(t *testing.T) {
	buf := &bytes.Buffer{}
	logger, err := New(buf, WarnLevel)
	assert.NoError(t, err)
	logger.Error("message")
	assert.True(t, strings.HasSuffix(buf.String(), "[ERR] message\n"))
}

func TestLogger_Errorf(t *testing.T) {
	buf := &bytes.Buffer{}
	logger, err := New(buf, DebugLevel)
	assert.NoError(t, err)
	logger.Errorf("message with %s", "arg")
	assert.True(t, strings.HasSuffix(buf.String(), "[ERR] message with arg\n"))
}
