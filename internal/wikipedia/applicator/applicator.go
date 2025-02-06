package applicator

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"goon-game/internal/wikipedia/config"
	"goon-game/internal/wikipedia/handlers"
	"goon-game/internal/wikipedia/infrastructure/cache"
	"goon-game/internal/wikipedia/infrastructure/message_brokers"
	"goon-game/internal/wikipedia/services"
	desc "goon-game/pkg/proto/wikipedia"
	"goon-game/pkg/utils"
	"net"
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
			cache.New,
			message_brokers.New,
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
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-shutdown
		in.Logger.Info("Stop server")
		ctx, cancel := context.WithTimeout(context.Background(), in.Cfg.ServerConfig.ShutdownTimeout)
		defer cancel()
		if err := in.Server.Shutdown(ctx); err != nil {
			in.Logger.Infof("Failure stop server: %v", err)
		}
	}()

	in.Logger.Infof("Start server on port: %s", in.Cfg.ServerConfig.Port)
	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterWikipediaServiceServer(s, in.Server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", in.Cfg.ServerConfig.Port))
	if err != nil {
		in.Logger.Fatalf("failed to listen: %v", err)
		return
	}

	go func() {
		if err = in.Server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			in.Logger.Fatalf("Failure start server: %v", err)
		}
	}()

	if err = s.Serve(lis); err != nil {
		in.Logger.Fatalf("Failure start server: %v", err)
	}
}
