package handlers

import (
	"context"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot"
	"goon-game/internal/discord_bot/config"
	"goon-game/pkg/utils"
)

type Server struct {
	cfg      *config.Config
	logger   utils.Logger
	services discord_bot.DiscordService
}

type ServerIn struct {
	fx.In
	Cfg      *config.Config
	Logger   utils.Logger
	Services discord_bot.DiscordService
}

func New(in ServerIn) *Server {
	return &Server{
		cfg:      in.Cfg,
		logger:   in.Logger,
		services: in.Services,
	}
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.services.Start(); err != nil {
		return err
	}

	s.services.InitHandlers()

	s.services.RetrieveWikipediaNotification(ctx)

	s.logger.Info("chezanah")

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.services.Stop()
}
