package message_brokers

import (
	"github.com/IBM/sarama"
	"go.uber.org/fx"
	"goon-game/internal/wikipedia/config"
	"goon-game/pkg/utils"
)

type kafkaBroker struct {
	cfg      *config.Config
	logger   utils.Logger
	producer sarama.SyncProducer
}

type KafkaBrokerIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in KafkaBrokerIn) MessageBrokers {
	provider, err := sarama.NewSyncProducer([]string{in.Cfg.KafkaConfig.KafkaHost}, nil)
	if err != nil {
		in.Logger.Fatalf("Error creating Kafka producer: %v", err)
	}

	return &kafkaBroker{
		cfg:      in.Cfg,
		logger:   in.Logger,
		producer: provider,
	}
}
