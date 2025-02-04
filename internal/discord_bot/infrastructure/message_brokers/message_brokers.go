package message_brokers

type MessageBrokers interface {
	RetrieveMessage() chan string
}
