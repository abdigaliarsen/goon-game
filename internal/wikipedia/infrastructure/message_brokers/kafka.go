package message_brokers

import (
	"github.com/IBM/sarama"
	"go.uber.org/fx"
	"goon-game/internal/wikipedia/config"
	"goon-game/pkg/utils"
	"time"
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
	conf := sarama.NewConfig()

	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = 5
	conf.Producer.Retry.Backoff = 500 * time.Millisecond
	conf.Producer.Return.Successes = true

	conf.Metadata.RefreshFrequency = 10 * time.Second

	conf.Version = sarama.V2_5_0_0

	provider, err := sarama.NewSyncProducer(
		[]string{in.Cfg.KafkaConfig.KafkaHost},
		conf,
	)
	if err != nil {
		in.Logger.Fatalf("Error creating Kafka producer: %v", err)
	}

	return &kafkaBroker{
		cfg:      in.Cfg,
		logger:   in.Logger,
		producer: provider,
	}
}
