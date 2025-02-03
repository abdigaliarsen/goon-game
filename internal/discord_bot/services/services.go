package services

import (
	"go.uber.org/fx"
	"goon-game/internal/discord_bot"
	"goon-game/internal/discord_bot/config"
	"goon-game/internal/discord_bot/transport/wikipedia"
	"goon-game/pkg/utils"
)

type discordService struct {
	cfg             *config.Config
	logger          utils.Logger
	wikipediaClient *wikipedia.GRPCTransport
}

type DiscordServiceIn struct {
	fx.In
	Cfg             *config.Config
	Logger          utils.Logger
	WikipediaClient *wikipedia.GRPCTransport
}

func New(in DiscordServiceIn) discord_bot.DiscordService {
	return &discordService{
		cfg:    in.Cfg,
		logger: in.Logger,
	}
}
