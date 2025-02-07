package handlers

import (
	"context"
	"go.uber.org/fx"
	"goon-game/internal/wikipedia"
	"goon-game/internal/wikipedia/config"
	desc "goon-game/pkg/proto/wikipedia"
	"goon-game/pkg/utils"
)

type Server struct {
	desc.UnimplementedWikipediaServiceServer
	cfg              *config.Config
	logger           utils.Logger
	wikipediaService wikipedia.WikipediaService
}

type ServerIn struct {
	fx.In
	Cfg              *config.Config
	Logger           utils.Logger
	WikipediaService wikipedia.WikipediaService
}

func New(in ServerIn) *Server {
	return &Server{
		cfg:              in.Cfg,
		logger:           in.Logger,
		wikipediaService: in.WikipediaService,
	}
}

func (s *Server) Start() {
	s.wikipediaService.StartService()
}

func (s *Server) Run() {
	go s.wikipediaService.RunService()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.wikipediaService.StopService()
	return nil
}
