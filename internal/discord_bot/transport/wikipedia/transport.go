package wikipedia

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"goon-game/internal/discord_bot/config"
	desc "goon-game/pkg/proto/wikipedia"
	"goon-game/pkg/utils"
)

type GRPCTransport struct {
	cfg    *config.Config
	logger utils.Logger
	desc.WikipediaServiceClient
}

type GRPCTransportIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func NewGRPCTransport(in GRPCTransportIn) (*GRPCTransport, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.NewClient(in.Cfg.ServerConfig.WikipediaTransportHost, opts...)
	if err != nil {
		in.Logger.Fatalf("Error connecting to wikipedia server: %v", err)
	}

	client := desc.NewWikipediaServiceClient(conn)
	return &GRPCTransport{
		cfg:                    in.Cfg,
		logger:                 in.Logger,
		WikipediaServiceClient: client,
	}, nil
}
