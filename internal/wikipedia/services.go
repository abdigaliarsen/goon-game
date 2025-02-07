package wikipedia

import (
	"goon-game/internal/wikipedia/dto"
	"time"
)

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
	GetLanguageUpdatesByDate(date time.Time) ([]*dto.LanguageUpdate, error)
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
	HardRestartService()
	RunService()
}
