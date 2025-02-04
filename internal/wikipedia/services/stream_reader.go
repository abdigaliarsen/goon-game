package services

import (
	"encoding/json"
	"goon-game/internal/wikipedia/dto"
	"io"
	"net/http"
	"regexp"
	"time"
)

var re = regexp.MustCompile(`lang="([^"]+)"`)

func (w *wikipediaService) ReadStream() chan dto.RecentChange {
	resp, err := http.Get(w.cfg.WikipediaConfig.StreamDataUrl)
	if err != nil {
		w.logger.Fatalf("Error fetching stream data: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.logger.Fatalf("Error fetching stream data: %d", resp.StatusCode)
	}

	changesChannel := make(chan dto.RecentChange)

	go func() {
		for w.running {
			buf := make([]byte, 2<<10)
			n, err := resp.Body.Read(buf)
			if err != nil && err != io.EOF {
				w.logger.Fatalf("Error reading stream data: %v", err)
			}

			if n == 0 {
				w.logger.Warnf("Error reading stream data: EOF")
				time.Sleep(2 * time.Second)
				continue
			}

			var change dto.RecentChange
			if err = json.Unmarshal(buf[:n], &change); err != nil {
				w.logger.Fatalf("Error reading stream data: %v", err)
			}

			if w.validLanguage(change) {
				changesChannel <- change
			}
		}
	}()

	return changesChannel
}

func (w *wikipediaService) validLanguage(change dto.RecentChange) bool {
	if w.language == "" {
		return true
	}

	matches := re.FindStringSubmatch(change.Data.ParsedComment)
	return len(matches) > 1 && w.language == matches[1]
}
