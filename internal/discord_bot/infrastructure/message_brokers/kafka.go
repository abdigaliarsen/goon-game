package message_brokers

import (
	"github.com/IBM/sarama"
	"go.uber.org/fx"
	"goon-game/internal/discord_bot/config"
	"goon-game/pkg/utils"
)

type kafkaBroker struct {
	cfg           *config.Config
	logger        utils.Logger
	consumerGroup sarama.ConsumerGroup
	consumer      *Consumer
}

type KafkaBrokerIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in KafkaBrokerIn) MessageBrokers {
	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{in.Cfg.KafkaConfig.KafkaHost},
		in.Cfg.KafkaConfig.KafkaWikipediaGroupID,
		nil)
	if err != nil {
		in.Logger.Fatal(err)
	}

	consumer := NewConsumer()

	return &kafkaBroker{
		cfg:           in.Cfg,
		logger:        in.Logger,
		consumerGroup: consumerGroup,
		consumer:      consumer,
	}
}
