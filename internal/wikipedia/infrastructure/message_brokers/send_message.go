package message_brokers

import "github.com/IBM/sarama"

func (k *kafkaBroker) SendMessage(message, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := k.producer.SendMessage(msg)
	if err != nil {
		k.logger.Errorf("Failed to send message to Kafka topic %s: %v", topic, err)
	} else {
		k.logger.Infof("Message sent to Kafka topic %s", topic)
	}

	return err
}
