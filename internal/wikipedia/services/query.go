package services

import (
	"errors"
	"goon-game/internal/wikipedia/dto"
	"goon-game/internal/wikipedia/utils"
	"time"
)

func (w *wikipediaService) SetLanguage(language string) error {
	if err := w.redis.AddS(utils.LanguageUpdatesKey, language); err != nil {
		return err
	}

	w.logger.Infof("Set language: %s", language)

	w.language = language

	w.HardRestartService()

	return w.redis.SetS(utils.LanguageKey, language)
}

func (w *wikipediaService) GetLanguage() (string, error) {
	w.logger.Infof("Get language: %s", w.language)
	return w.redis.GetS(utils.LanguageKey)
}

func (w *wikipediaService) GetLanguageUpdates() ([]*dto.LanguageUpdate, error) {
	languages, updatedAt, err := w.redis.GetList(utils.LanguageUpdatesKey)
	if err != nil {
		return nil, err
	}

	if len(languages) != len(updatedAt) {
		return nil, errors.New("language updates not found")
	}

	languageUpdates := make([]*dto.LanguageUpdate, 0, len(languages))
	for i := range languages {
		languageUpdates = append(languageUpdates, &dto.LanguageUpdate{
			Language:  languages[i],
			UpdatedAt: time.Unix(updatedAt[i], 0),
		})
	}

	w.logger.Infof("Get language updates: %v", languageUpdates)

	return languageUpdates, nil
}

func (w *wikipediaService) GetLanguageUpdatesByDate(date time.Time) ([]*dto.LanguageUpdate, error) {
	languages, timestamps, err := w.redis.GetZRangeByDate(date, utils.LanguageUpdatesKey)
	if err != nil {
		return nil, err
	}

	if len(languages) != len(timestamps) {
		return nil, errors.New("language updates not found")
	}

	languageUpdates := make([]*dto.LanguageUpdate, 0, len(languages))
	for i := range languages {
		languageUpdates = append(languageUpdates, &dto.LanguageUpdate{
			Language:  languages[i],
			UpdatedAt: time.Unix(timestamps[i], 0),
		})
	}

	w.logger.Infof("Get language updates: %v", languageUpdates)
	return languageUpdates, nil
}
