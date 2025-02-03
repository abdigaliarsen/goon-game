package handlers

import (
	"context"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/config"
	"goon-game/pkg/utils"
)

type Server struct {
	cfg    *config.Config
	logger utils.Logger
}

type ServerIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in ServerIn) *Server {
	return &Server{
		cfg:    in.Cfg,
		logger: in.Logger,
	}
}

func (s *Server) Start() error {

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {

	return nil
}
