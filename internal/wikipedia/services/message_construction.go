package services

import (
	"errors"
	"fmt"
	"goon-game/internal/wikipedia/dto"
)

func (w *wikipediaService) ConstructMessage(recentChange *dto.RecentChange) (string, error) {
	if recentChange == nil {
		return "", errors.New("recentChange is nil")
	}

	msg := fmt.Sprintf(
		"(%s, %s, %s, %d)",
		recentChange.Data.Title,
		recentChange.Data.TitleURL,
		recentChange.Data.User,
		recentChange.Data.Timestamp,
	)

	return msg, nil
}
