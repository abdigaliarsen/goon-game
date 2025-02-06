package services

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot"
	"goon-game/internal/discord_bot/config"
	"goon-game/internal/discord_bot/infrastructure/cache"
	"goon-game/internal/discord_bot/infrastructure/message_brokers"
	"goon-game/internal/discord_bot/transport/wikipedia"
	"goon-game/pkg/utils"
)

type discordService struct {
	cfg             *config.Config
	logger          utils.Logger
	wikipediaClient *wikipedia.GRPCTransport
	discord         *discordgo.Session
	kafka           message_brokers.MessageBrokers
	cache           cache.Cache

	running bool
}

type DiscordServiceIn struct {
	fx.In
	Cfg             *config.Config
	Logger          utils.Logger
	WikipediaClient *wikipedia.GRPCTransport
	Kafka           message_brokers.MessageBrokers
	Cache           cache.Cache
}

func New(in DiscordServiceIn) (discord_bot.DiscordService, error) {
	discord, err := discordgo.New("Bot " + in.Cfg.DiscordApiConfig.DiscordApiToken)
	if err != nil {
		return nil, err
	}

	return &discordService{
		cfg:             in.Cfg,
		logger:          in.Logger,
		wikipediaClient: in.WikipediaClient,
		discord:         discord,
		kafka:           in.Kafka,
		cache:           in.Cache,
		running:         false,
	}, nil
}
