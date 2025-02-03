package wikipedia

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"goon-game/internal/discord_bot/config"
	desc "goon-game/pkg/proto/wikipedia"
)

type GRPCTransport struct {
	cfg *config.Config
	desc.WikipediaServiceClient
}

type GRPCTransportIn struct {
	fx.In
	Cfg *config.Config
}

func NewGRPCTransport(in GRPCTransportIn) (*GRPCTransport, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.NewClient(in.Cfg.ServerConfig.WikipediaTransportHost, opts...)

	if err != nil {
		return nil, err
	}

	client := desc.NewWikipediaServiceClient(conn)
	return &GRPCTransport{
		cfg:                    in.Cfg,
		WikipediaServiceClient: client,
	}, nil
}
