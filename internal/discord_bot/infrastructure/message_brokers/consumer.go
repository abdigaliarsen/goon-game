package message_brokers

import (
	"github.com/IBM/sarama"
)

type ConsumerHandler struct {
	message chan string
}

func NewConsumerHandler() *ConsumerHandler {
	return &ConsumerHandler{message: make(chan string, 100)}
}

func (c *ConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	close(c.message)
	return nil
}

func (c *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		c.message <- string(msg.Value)
		session.MarkMessage(msg, "")
	}

	return nil
}
