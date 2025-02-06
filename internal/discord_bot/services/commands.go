package services

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	desc "goon-game/pkg/proto/wikipedia"
	"strings"
)

func (d *discordService) SetLanguage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var language string
	if i.ApplicationCommandData().Options != nil && len(i.ApplicationCommandData().Options) > 0 {
		language = i.ApplicationCommandData().Options[0].StringValue()
	}

	d.logger.Infof("Set language: %s", language)

	_, err := d.wikipediaClient.SetLanguage(context.TODO(), &desc.SetLanguageRequest{Language: language})
	if err != nil {
		d.sendMessage(s, i, fmt.Sprintf("Error setting language: %v", err))
		return
	}

	d.logger.Infof("Setting language: %v", language)
	d.sendMessage(s, i, fmt.Sprintf("Setting language success: %v", language))
}

func (d *discordService) GetLanguageUpdates(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp, err := d.wikipediaClient.GetLanguageUpdates(context.TODO(), &desc.EmptyRequest{})
	if err != nil {
		d.logger.Errorf("Error getting language updates: %v", err)
		d.sendMessage(s, i, fmt.Sprintf("Error getting language updates: %v", err))
		return
	}

	languageUpdates := make([]string, 0, len(resp.GetUpdates()))
	for _, update := range resp.GetUpdates() {
		updateLanguage := "all"
		if update.Language != "" {
			updateLanguage = update.Language
		}

		languageUpdates = append(languageUpdates, fmt.Sprintf("%s: %s", updateLanguage, update.UpdatedAt.AsTime().String()))
	}

	message := strings.Join(languageUpdates, "\n")
	d.sendMessage(s, i, message)

	d.logger.Infof("Getting language updates: %v", languageUpdates)
}

func (d *discordService) sendMessage(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	}); err != nil {
		d.logger.Fatalf("Error sending message: %v", err)
	}
}
