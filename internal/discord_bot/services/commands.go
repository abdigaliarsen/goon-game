package services

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goon-game/internal/discord_bot/utils"
	desc "goon-game/pkg/proto/wikipedia"
	"strings"
)

func (d *discordService) SetLanguage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optMap := utils.ParseOptions(i.ApplicationCommandData().Options)
	language := optMap["message"].StringValue()

	_, err := d.wikipediaClient.SetLanguage(context.TODO(), &desc.SetLanguageRequest{Language: language})
	if err != nil {
		d.logger.Errorf("Error setting language: %v", err)
		if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Error setting language: %v", err),
			},
		}); err != nil {
			d.logger.Fatalf("Error setting language: %v", err)
		}

		return
	}

	d.logger.Infof("Setting language: %v", language)
	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Setting language success: %v", language),
		},
	}); err != nil {
		d.logger.Fatalf("Error setting language: %v", err)
	}
}

func (d *discordService) GetLanguageUpdates(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp, err := d.wikipediaClient.GetLanguageUpdates(context.TODO(), &desc.EmptyRequest{})
	if err != nil {
		d.logger.Errorf("Error getting language updates: %v", err)
		return
	}

	languageUpdates := make([]string, 0, len(resp.GetUpdates()))
	for _, update := range resp.GetUpdates() {
		languageUpdates = append(languageUpdates, fmt.Sprintf("%s: %s", update.Language, update.UpdatedAt.AsTime().String()))
	}

	message := strings.Join(languageUpdates, "\n")
	if err = d.discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	}); err != nil {
		d.logger.Fatalf("Error setting language updates: %v", err)
	}

	d.logger.Infof("Getting language updates: %v", languageUpdates)
}
