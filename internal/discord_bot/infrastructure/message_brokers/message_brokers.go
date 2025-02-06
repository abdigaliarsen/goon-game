package message_brokers

import "context"

type MessageBrokers interface {
	RetrieveMessage(ctx context.Context) chan string
}
