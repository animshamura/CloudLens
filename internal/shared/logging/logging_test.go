package logging

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLoggerReturnsConfiguredLogger(t *testing.T) {
	logger := NewLogger("development")
	require.NotNil(t, logger)
}
