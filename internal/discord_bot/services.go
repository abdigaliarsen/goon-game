package discord_bot

type DiscordService interface {
	SendMessageService
}

type SendMessageService interface {
	SendMessage() error
}
