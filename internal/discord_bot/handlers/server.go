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
	services discord_bot.DiscordService
}

func New(in ServerIn) *Server {
	return &Server{
		cfg:    in.Cfg,
		logger: in.Logger,
	}
}

func (s *Server) Start() error {
	s.services.InitHandlers()

	if err := s.services.Start(); err != nil {
		return err
	}

	go s.services.RetrieveWikipediaNotification()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.services.Stop()
}
