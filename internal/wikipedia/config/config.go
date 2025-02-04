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
	KafkaConfig     KafkaConfig
	RedisConfig     RedisConfig
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

type KafkaConfig struct {
	KafkaHost string `env:"KAFKA_HOST"`
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
		RedisConfig: RedisConfig{
			Addr:     utils.MustGetEnv[string]("REDIS_ADDR"),
			Password: utils.MustGetEnv[string]("REDIS_PASSWORD"),
			DB:       utils.MustGetEnv[int]("REDIS_DB"),
		},
		KafkaConfig: KafkaConfig{
			KafkaHost: utils.MustGetEnv[string]("KAFKA_HOST"),
		},
		ServerConfig: ServerConfig{
			Port:            utils.MustGetEnv[string]("PORT"),
			ShutdownTimeout: utils.MustGetEnv[time.Duration]("SHUTDOWN_TIMEOUT"),
		},
		WikipediaConfig: WikipediaConfig{
			StreamDataUrl: utils.MustGetEnv[string]("STREAM_DATA_URL"),
		},
		LogConfig: LogConfig{
			ENV: utils.MustGetEnv[string]("ENV"),
		},
	}

	return cfg, nil
}
