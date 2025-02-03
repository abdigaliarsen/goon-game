package config

import (
	"github.com/lpernett/godotenv"
	"goon-game/pkg/utils"
	"os"
)

type Config struct {
	LogConfig        LogConfig
	DiscordApiConfig DiscordApiConfig
}

type DiscordApiConfig struct {
	Token string `env:"TOKEN"`
}

type LogConfig struct {
	ENV string `env:"ENV"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	cfg := &Config{
		DiscordApiConfig: DiscordApiConfig{
			Token: os.Getenv("DISCORD_API_TOKEN"),
		},
		LogConfig: LogConfig{
			ENV: utils.MustGetEnv[string]("ENV"),
		},
	}

	return cfg, nil
}
