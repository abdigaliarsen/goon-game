package main

import (
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/applicator"
	"goon-game/internal/discord_bot/config"
	"goon-game/internal/discord_bot/infrastructure/logger"
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
