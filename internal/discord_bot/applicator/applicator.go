package applicator

import (
	"context"
	"errors"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/config"
	"goon-game/internal/discord_bot/handlers"
	"goon-game/internal/discord_bot/infrastructure/cache"
	"goon-game/internal/discord_bot/infrastructure/message_brokers"
	"goon-game/internal/discord_bot/services"
	"goon-game/internal/discord_bot/transport/wikipedia"
	"goon-game/pkg/utils"
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
			message_brokers.New,
			cache.New,
			wikipedia.NewGRPCTransport,
			services.New,
			handlers.New,
		),
		fx.Invoke(run),
	)

	app.Run()
}

type deps struct {
	fx.In
	Server *handlers.Server
	Logger utils.Logger
	Cfg    *config.Config
}

func run(in deps) {
	ctx := context.Background()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-shutdown
		in.Logger.Info("Stop server")
		ctx, cancel := context.WithTimeout(ctx, in.Cfg.ServerConfig.ShutdownTimeout)
		defer cancel()
		if err := in.Server.Shutdown(ctx); err != nil {
			in.Logger.Infof("Failure stop server: %v", err)
		}
	}()

	in.Logger.Infof("Start server on port: %s", in.Cfg.ServerConfig.Port)
	if err := in.Server.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		in.Logger.Fatalf("Failure start server: %v", err)
	}

	in.Logger.Info("chezanah")
}
