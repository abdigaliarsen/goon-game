package services

import (
	"context"
	"goon-game/internal/discord_bot/utils"
)

func (d *discordService) RetrieveWikipediaNotification(ctx context.Context) {
	d.logger.Info("Start retrieving wikipedia notifications")

	messages := d.kafka.RetrieveMessage(ctx)

	for message := range messages {
		if !d.running {
			d.logger.Info("Discord service is not running")
			return
		}

		discordChannelIds, _, err := d.cache.GetArr(utils.DiscordChannelIdsKey)
		if err != nil {
			d.logger.Errorf("Error getting discord channel ids: %v", err)
			continue
		}

		for _, channelId := range discordChannelIds {
			_, err = d.discord.ChannelMessageSend(channelId, message)
			if err != nil {
				d.logger.Errorf("Error sending message to discord channel: %v", err)
				continue
			}
		}
	}

	d.logger.Info("Done retrieving wikipedia notifications")
}
