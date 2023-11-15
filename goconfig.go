package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Config interface{}

func ConfigFromFile(path string, cfg Config) error {
	f, err := os.OpenFile(path, os.O_RDONLY | os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".yaml", ".yml":
		err = ParseYAML(f, cfg)
	case ".json":
		err = ParseJSON(f, cfg)
	}
	return err
}

func ParseYAML(r io.Reader, cfg Config) error {
	return yaml.NewDecoder(r).Decode(cfg)
}

func ParseJSON(r io.Reader, cfg Config) error {
	return json.NewDecoder(r).Decode(cfg)
}

type ServerConfig struct {
	Host         string `yaml:host`
	Port         string `yaml:port`
	ReadTimeout  time.Duration `yaml:readtimeout`
	WriteTimeout time.Duration `yaml:writetimeout`
}

func ServerConfigFromYaml(filepath string) (cfg ServerConfig, err error) {
	err = ConfigFromFile(filepath, &cfg)
	return cfg, err
}