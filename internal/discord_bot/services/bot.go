package services

import (
	"github.com/bwmarrin/discordgo"
	"goon-game/internal/discord_bot/utils"
)

func (s *discordService) Start() error {
	s.running = true
	return s.discord.Open()
}

func (s *discordService) Stop() error {
	s.running = false
	return s.discord.Close()
}

func (s *discordService) InitHandlers() {
	descriptors := s.getDescriptors()

	for _, desc := range descriptors {
		s.discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if desc.Predicate(i.ApplicationCommandData().Name, i.Type) {
				desc.Method(s, i)
			}
		})
	}
}

func (s *discordService) getDescriptors() []utils.MethodDescriptor {
	return []utils.MethodDescriptor{
		{
			Predicate: utils.CommandEqual("!setLang", discordgo.InteractionApplicationCommand),
			Method:    s.SetLanguage,
		},
		{
			Predicate: utils.CommandEqual("!recent", discordgo.InteractionApplicationCommand),
			Method:    s.GetLanguageUpdates,
		},
	}
}
