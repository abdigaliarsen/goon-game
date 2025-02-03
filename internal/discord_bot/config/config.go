package config

import (
	"github.com/lpernett/godotenv"
	"goon-game/pkg/utils"
	"os"
	"time"
)

type Config struct {
	LogConfig        LogConfig
	DiscordApiConfig DiscordApiConfig
	ServerConfig     ServerConfig
}

type ServerConfig struct {
	WikipediaTransportHost string        `env:"WIKIPEDIA_TRANSPORT_HOST"`
	Port                   string        `env:"PORT"`
	ShutdownTimeout        time.Duration `env:"SHUTDOWN_TIMEOUT"`
}

type DiscordApiConfig struct {
	DiscordApiToken string `env:"DISCORD_API_TOKEN"`
}

type LogConfig struct {
	ENV string `env:"ENV"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	cfg := &Config{
		ServerConfig: ServerConfig{
			WikipediaTransportHost: utils.MustGetEnv[string]("WIKIPEDIA_TRANSPORT_HOST"),
			Port:                   utils.MustGetEnv[string]("PORT"),
			ShutdownTimeout:        utils.MustGetEnv[time.Duration]("SHUTDOWN_TIMEOUT"),
		},
		DiscordApiConfig: DiscordApiConfig{
			DiscordApiToken: utils.MustGetEnv[string]("DISCORD_API_TOKEN"),
		},
		LogConfig: LogConfig{
			ENV: utils.MustGetEnv[string]("ENV"),
		},
	}

	return cfg, nil
}
