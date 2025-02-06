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
		k.logger.Infof("Message \"%s\" sent to Kafka topic %s", message, topic)
	}

	return err
}
