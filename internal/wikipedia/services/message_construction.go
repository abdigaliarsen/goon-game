package services

import (
	"errors"
	"fmt"
	"goon-game/internal/wikipedia/dto"
	"time"
)

func (w *wikipediaService) ConstructMessage(recentChange *dto.RecentChange) (string, error) {
	if recentChange == nil {
		return "", errors.New("recentChange is nil")
	}

	msg := fmt.Sprintf(
		"(%s, %s, %s, %s)",
		recentChange.Data.Title,
		recentChange.Data.TitleURL,
		recentChange.Data.User,
		time.Unix(recentChange.Data.Timestamp, 0).UTC().Format(time.RFC3339),
	)

	return msg, nil
}
