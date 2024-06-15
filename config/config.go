package config

import (
	"fmt"
	"os"

	"github.com/inview-team/sadko_indexer/internal/infrastructure/tag_service/rabbitmq"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/video_repository/postgres"
	"gopkg.in/yaml.v3"
)

type Config struct {
	PostgresConfig postgres.Config `yaml:"postgres"`
	RabbitConfig   rabbitmq.Config `yaml:"rabbit"`
}

func Load(s string) (*Config, error) {
	cfg := &Config{}

	err := yaml.Unmarshal([]byte(s), cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadFile(filename string) (*Config, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to find configuration file")
	}
	cfg, err := Load(string(content))
	if err != nil {
		return nil, fmt.Errorf("parsing YAML file %s failed: %v", filename, err)
	}
	return cfg, nil
}
