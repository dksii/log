package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItoa(t *testing.T) {
	ttable := []struct {
		i        int
		wid      int
		expected []byte
	}{
		{i: 123, wid: 2, expected: []byte("123")},
		{i: 123, wid: 3, expected: []byte("123")},
		{i: 123, wid: 4, expected: []byte("0123")},
	}

	for _, tt := range ttable {
		var buf []byte
		itoa(&buf, tt.i, tt.wid)
		assert.Equal(t, tt.expected, buf)
	}
}
