package wikipedia

import "goon-game/internal/wikipedia/dto"

type WikipediaService interface {
	QueryService
	NotificationService
	StreamReaderService
	RunningStatusService
	MessageConstructionService
}

type QueryService interface {
	SetLanguage(language string) error
	GetLanguage() (string, error)
	GetLanguageUpdates() ([]*dto.LanguageUpdate, error)
}

type NotificationService interface {
	SendNotification(messageContent string) error
}

type MessageConstructionService interface {
	ConstructMessage(recentChange *dto.RecentChange) (string, error)
}

type StreamReaderService interface {
	ReadStream() chan dto.RecentChange
}

type RunningStatusService interface {
	StartService()
	StopService()
	RestartService()
}
