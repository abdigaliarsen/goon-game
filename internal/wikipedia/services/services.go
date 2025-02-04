package services

import (
	"go.uber.org/fx"
	"goon-game/internal/wikipedia"
	"goon-game/internal/wikipedia/config"
	"goon-game/internal/wikipedia/infrastructure/cache"
	"goon-game/internal/wikipedia/infrastructure/message_brokers"
	"goon-game/pkg/utils"
)

type wikipediaService struct {
	cfg    *config.Config
	logger utils.Logger
	kafka  message_brokers.MessageBrokers
	redis  cache.Cache

	running  bool
	language string
}

type WikipediaServiceIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
	Kafka  message_brokers.MessageBrokers
	Redis  cache.Cache
}

func New(in WikipediaServiceIn) wikipedia.WikipediaService {
	return &wikipediaService{
		cfg:     in.Cfg,
		logger:  in.Logger,
		redis:   in.Redis,
		kafka:   in.Kafka,
		running: true,
	}
}
