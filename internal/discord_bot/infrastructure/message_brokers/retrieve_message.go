package message_brokers

import (
	"context"
	"goon-game/pkg/utils"
)

func (k *kafkaBroker) RetrieveMessage() chan string {
	messages := make(chan string)

	go func() {
		for {
			if k.consumerGroup == nil {
				break
			}

			err := k.consumerGroup.Consume(context.TODO(), []string{utils.WikipediaTopic}, k.consumer)
			if err != nil {
				k.logger.Errorf("Error consuming messages: %v", err)
				continue
			}

			for msg := range k.consumer.message {
				messages <- msg
			}
		}

		close(messages)
	}()

	return messages
}
