package message_brokers

type MessageBrokers interface {
	SendMessage(message, topic string) error
}
