package services

import (
	"goon-game/pkg/utils"
)

func (w *wikipediaService) SendNotification(messageContent string) error {
	return w.kafka.SendMessage(messageContent, utils.WikipediaTopic)
}
