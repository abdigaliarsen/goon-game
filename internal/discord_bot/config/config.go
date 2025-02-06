package config

import (
	"github.com/lpernett/godotenv"
	"goon-game/pkg/utils"
	"os"
	"strings"
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
	DiscordApiToken       string   `env:"DISCORD_API_TOKEN"`
	DiscordApplicationId  string   `env:"DISCORD_APPLICATION_ID"`
	DiscordPublicKey      string   `env:"DISCORD_PUBLIC_KEY"`
	DiscordDefaultChatIds []string `env:"DISCORD_DEFAULT_CHAT_IDS"`
}

type LogConfig struct {
	ENV string `env:"ENV"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	discordChatIds := utils.MustGetEnv[string]("DISCORD_DEFAULT_CHAT_IDS")

	cfg := &Config{
		RedisConfig: RedisConfig{
			Addr:     utils.MustGetEnv[string]("REDIS_ADDR"),
			Password: utils.MustGetEnv[string]("REDIS_PASSWORD"),
			DB:       utils.MustGetEnv[int]("REDIS_DB"),
		},
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
			DiscordApiToken:       utils.MustGetEnv[string]("DISCORD_API_TOKEN"),
			DiscordApplicationId:  utils.MustGetEnv[string]("DISCORD_APPLICATION_ID"),
			DiscordPublicKey:      utils.MustGetEnv[string]("DISCORD_PUBLIC_KEY"),
			DiscordDefaultChatIds: strings.Split(discordChatIds, ","),
		},
		LogConfig: LogConfig{
			ENV: utils.MustGetEnv[string]("ENV"),
		},
	}

	return cfg, nil
}
