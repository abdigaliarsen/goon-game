package services

import (
	"go.uber.org/fx"
	"goon-game/internal/wikipedia"
	"goon-game/internal/wikipedia/config"
	"goon-game/pkg/utils"
)

type wikipediaService struct {
	cfg    *config.Config
	logger utils.Logger
}

type WikipediaServiceIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in WikipediaServiceIn) wikipedia.WikipediaService {
	return &wikipediaService{
		cfg:    in.Cfg,
		logger: in.Logger,
	}
}
