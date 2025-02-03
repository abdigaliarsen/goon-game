package config

import (
	"github.com/lpernett/godotenv"
	"goon-game/pkg/utils"
	"os"
	"time"
)

type Config struct {
	LogConfig       LogConfig
	WikipediaConfig WikipediaConfig
	ServerConfig    ServerConfig
}

type ServerConfig struct {
	Port            string        `env:"PORT"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT"`
}

type WikipediaConfig struct {
	StreamDataUrl string `env:"STREAM_DATA_URL"`
}

type LogConfig struct {
	ENV string `env:"ENV"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	cfg := &Config{
		WikipediaConfig: WikipediaConfig{
			StreamDataUrl: os.Getenv("STREAM_DATA_URL"),
		},
		LogConfig: LogConfig{
			ENV: utils.MustGetEnv[string]("ENV"),
		},
	}

	return cfg, nil
}
