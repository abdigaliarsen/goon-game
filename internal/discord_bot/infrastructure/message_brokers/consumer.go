package message_brokers

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type Consumer struct {
	message chan string
}

func NewConsumer() *Consumer {
	return &Consumer{message: make(chan string)}
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	close(c.message)
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var parsedMessage string
		if err := json.Unmarshal(msg.Value, &parsedMessage); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		c.message <- parsedMessage

		session.MarkMessage(msg, "")
	}

	return nil
}
