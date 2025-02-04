package handlers

import (
	"context"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/config"
	"goon-game/internal/wikipedia"
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

func (s *Server) Start() error {
	go func() {
		for change := range s.wikipediaService.ReadStream() {
			msg, err := s.wikipediaService.ConstructMessage(&change)
			if err != nil {
				s.logger.Errorf("Error constructing message: %v", err)
				continue
			}

			if msg != "" {
				if err := s.wikipediaService.SendNotification(msg); err != nil {
					s.logger.Errorf("Error sending notification: %v", err)
				}
			}
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.wikipediaService.Stop()
	return nil
}
