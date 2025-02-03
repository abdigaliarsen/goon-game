package applicator

import (
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/config"
	"goon-game/pkg/utils"
)

type Applicator struct {
	cfg    *config.Config
	logger utils.Logger
}

type ApplicatorIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in ApplicatorIn) *Applicator {
	return &Applicator{
		cfg:    in.Cfg,
		logger: in.Logger,
	}
}

func (a *Applicator) Run() {
	app := fx.New(
		fx.Provide(),
	)
	app.Run()
}
