package services

import "goon-game/internal/discord_bot/utils"

func (d *discordService) RetrieveWikipediaNotification() {
	for {
		if !d.running {
			break
		}

		for message := range d.kafka.RetrieveMessage() {
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
	}
}
