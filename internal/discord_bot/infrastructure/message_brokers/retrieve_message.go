package message_brokers

import (
	"context"
	"github.com/IBM/sarama"
	"goon-game/pkg/utils"
	"time"
)

func (k *kafkaBroker) RetrieveMessage(ctx context.Context) chan string {
	messages := make(chan string, 100)
	k.logger.Info("Created message retrieving channel")

	go func() {
		for {
			k.logger.Info("Waiting for messages")

			if k.consumer == nil {
				k.logger.Errorf("Consumer group is nil")
				break
			}

			partitionConsumer, err := k.consumer.ConsumePartition(utils.WikipediaTopic, 0, sarama.OffsetNewest)
			if err != nil {
				k.logger.Errorf("Error consuming messages: %v", err)
				time.Sleep(2 * time.Second)
				continue
			}

			for message := range partitionConsumer.Messages() {
				k.logger.Infof("Received message: %s", string(message.Value))
				messages <- string(message.Value)
			}
		}
	}()

	return messages
}
