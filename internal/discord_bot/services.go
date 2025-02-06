package discord_bot

import (
	"context"
	"github.com/bwmarrin/discordgo"
)

type DiscordService interface {
	BotService
	CommandsService
	WikipediaNotificationRetrieverService
}

type BotService interface {
	Start() error
	Stop() error
	InitHandlers()
}

type CommandsService interface {
	SetLanguage(s *discordgo.Session, i *discordgo.InteractionCreate)
	GetLanguageUpdates(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type WikipediaNotificationRetrieverService interface {
	RetrieveWikipediaNotification(ctx context.Context)
}
