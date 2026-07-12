package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadSetsDefaultsAndEnvironmentOverrides(t *testing.T) {
	t.Setenv("HTTP_PORT", "9090")
	t.Setenv("APP_ENV", "production")

	cfg, err := Load()
	require.NoError(t, err)
	require.Equal(t, "9090", cfg.HTTP.Port)
	require.Equal(t, "production", cfg.App.Env)
	require.Equal(t, "0.0.0.0", cfg.HTTP.Host)
}
