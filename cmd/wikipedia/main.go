package main

import (
	"go.uber.org/fx"
	"goon-game/internal/wikipedia/applicator"
	"goon-game/internal/wikipedia/config"
	"goon-game/internal/wikipedia/infrastructure/logger"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			logger.New,
			applicator.New,
		),
		fx.Invoke(
			func(app *applicator.Applicator) {
				app.Run()
			},
		),
	)

	app.Run()
}
