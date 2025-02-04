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
	KafkaConfig      KafkaConfig
	RedisConfig      RedisConfig
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

type KafkaConfig struct {
	KafkaHost             string `env:"KAFKA_HOST"`
	KafkaWikipediaGroupID string `env:"KAFKA_WIKIPEDIA_GROUP_ID"`
}

type ServerConfig struct {
	WikipediaTransportHost string        `env:"WIKIPEDIA_TRANSPORT_HOST"`
	Port                   string        `env:"PORT"`
	ShutdownTimeout        time.Duration `env:"SHUTDOWN_TIMEOUT"`
}

type DiscordApiConfig struct {
	DiscordApiToken      string `env:"DISCORD_API_TOKEN"`
	DiscordApplicationId string `env:"DISCORD_APPLICATION_ID"`
	DiscordPublicKey     string `env:"DISCORD_PUBLIC_KEY"`
}

type LogConfig struct {
	ENV string `env:"ENV"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	cfg := &Config{
		KafkaConfig: KafkaConfig{
			KafkaHost:             utils.MustGetEnv[string]("KAFKA_HOST"),
			KafkaWikipediaGroupID: utils.MustGetEnv[string]("KAFKA_WIKIPEDIA_GROUP_ID"),
		},
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
