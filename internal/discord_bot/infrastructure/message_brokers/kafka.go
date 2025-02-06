package message_brokers

import (
	"github.com/IBM/sarama"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/config"
	"goon-game/pkg/utils"
	"time"
)

type kafkaBroker struct {
	cfg             *config.Config
	logger          utils.Logger
	consumer        sarama.Consumer
	consumerHandler *ConsumerHandler
}

type KafkaBrokerIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in KafkaBrokerIn) MessageBrokers {
	conf := sarama.NewConfig()

	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	conf.Consumer.Group.Session.Timeout = 30 * time.Second
	conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
	}

	conf.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	conf.Consumer.Group.Rebalance.Timeout = 60 * time.Second

	conf.Metadata.RefreshFrequency = 10 * time.Second

	conf.Version = sarama.V2_5_0_0

	consumer, err := sarama.NewConsumer(
		[]string{in.Cfg.KafkaConfig.KafkaHost},
		conf,
	)

	if err != nil {
		in.Logger.Fatal(err)
	}

	consumerHandler := NewConsumerHandler()

	return &kafkaBroker{
		cfg:             in.Cfg,
		logger:          in.Logger,
		consumer:        consumer,
		consumerHandler: consumerHandler,
	}
}
