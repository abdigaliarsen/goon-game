package services

import (
	"github.com/bwmarrin/discordgo"
	"goon-game/internal/discord_bot/utils"
)

func (s *discordService) Start() error {
	if s.running {
		s.logger.Warn("Bot is already running")
		return nil
	}

	s.logger.Info("Starting Discord bot")

	s.running = true

	for _, chatId := range s.cfg.DiscordApiConfig.DiscordDefaultChatIds {
		if err := s.cache.Add(utils.DiscordChannelIdsKey, chatId); err != nil {
			return err
		}
	}

	return s.discord.Open()
}

func (s *discordService) Stop() error {
	s.logger.Info("Shutting down Discord bot")

	s.running = false
	return s.discord.Close()
}

func (s *discordService) InitHandlers() {
	handlers := s.getHandlers()
	s.discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	commands := s.getCommands()
	for _, v := range commands {
		s.logger.Infof("Adding '%s' command", v.Name)

		_, err := s.discord.ApplicationCommandCreate(s.discord.State.User.ID, "", v)
		if err != nil {
			s.logger.Fatalf("Error creating command: %+v", err)
		}
	}
}

func (s *discordService) getHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"setlang": s.SetLanguage,
		"recent":  s.GetLanguageUpdates,
		"stats":   s.GetStats,
	}
}

func (s *discordService) getCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "setlang",
			Description: "Set the Wikipedia language",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "language",
					Description: "Language to filter recent changes",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "recent",
			Description: "Get recent language updates",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "stats",
			Description: "Track the number of changes per day for each language.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "date",
					Description: "Show number of changes on specified date [yyyy-mm-dd]",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
	}
}
