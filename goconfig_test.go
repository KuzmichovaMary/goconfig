package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerConfigFromFile(t *testing.T) {
	filepath := "config.yaml"
	cfg_, _ := ServerConfigFromYaml(filepath)
	cfg := ServerConfig{
		"localhost",
		"8080",
		10 * time.Nanosecond,
		10 * time.Nanosecond,
	}
	assert.Equal(t, cfg, cfg_)
}