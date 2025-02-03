package applicator

import (
	"context"
	"errors"
	"go.uber.org/fx"
	"goon-game/internal/wikipedia/config"
	"goon-game/internal/wikipedia/handlers"
	"goon-game/internal/wikipedia/services"
	"goon-game/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		fx.Provide(
			func() utils.Logger { return a.logger },
			func() *config.Config { return a.cfg },
			services.New,
			handlers.New,
		),
		fx.Invoke(
			func(server *handlers.Server) {
				shutdown := make(chan os.Signal, 1)
				signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
				go func() {
					<-shutdown
					a.logger.Info("Stop server")
					ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ServerConfig.ShutdownTimeout)
					defer cancel()
					if err := server.Shutdown(ctx); err != nil {
						a.logger.Infof("Failure stop server: %v", err)
					}
				}()

				a.logger.Infof("Start server on port: %s", a.cfg.ServerConfig.Port)
				if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("Failure start server: %v", err)
				}
			},
		),
	)

	app.Run()
}
