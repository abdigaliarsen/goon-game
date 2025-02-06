package services

import (
	"bufio"
	"encoding/json"
	"goon-game/internal/wikipedia/dto"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var re = regexp.MustCompile(`lang="([^"]+)"`)

func (w *wikipediaService) ReadStream() chan dto.RecentChange {
	changesChannel := make(chan dto.RecentChange, 100)

	go func() {
		for w.running {
			resp := w.doRecentChange()
			if resp == nil {
				w.logger.Warn("Failed to establish connection, retrying in 5 seconds...")
				time.Sleep(5 * time.Second)
				continue
			}

			w.processStream(resp, changesChannel)

			time.Sleep(500 * time.Millisecond)
		}
	}()

	return changesChannel
}

func (w *wikipediaService) processStream(resp *http.Response, changesChannel chan dto.RecentChange) {
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	timeout := time.NewTimer(5 * time.Second)

	var (
		eventType string
		idData    strings.Builder
		jsonData  strings.Builder
	)

	for w.running {
		select {
		case <-timeout.C:
			w.logger.Warn("No response for 5 seconds, reconnecting...")
			return

		default:
			for scanner.Scan() {
				timeout.Reset(5 * time.Second)

				line := scanner.Text()
				if line == "" {
					w.processEvent(eventType, idData.String(), jsonData.String(), changesChannel)

					eventType = ""
					idData.Reset()
					jsonData.Reset()
					continue
				}

				if strings.HasPrefix(line, "event: ") {
					eventType = strings.TrimPrefix(line, "event: ")
				} else if strings.HasPrefix(line, "id: ") {
					idData.WriteString(strings.TrimPrefix(line, "id: "))
				} else if strings.HasPrefix(line, "data: ") {
					jsonData.WriteString(strings.TrimPrefix(line, "data: "))
				}
			}

			w.logger.Warn("Stream closed by server, reconnecting...")
			return
		}
	}

	if err := scanner.Err(); err != nil {
		w.logger.Errorf("Error reading stream data: %v", err)
	}
}

func (w *wikipediaService) processEvent(eventType, idData, jsonData string, changesChannel chan dto.RecentChange) {
	if jsonData == "" {
		return
	}

	var change dto.RecentChange

	if idData != "" {
		if err := json.Unmarshal([]byte(idData), &change.Id); err != nil {
			w.logger.Errorf("Error parsing Id field: %v. Raw ID Data: %s", err, idData)
			return
		}
	}

	if err := json.Unmarshal([]byte(jsonData), &change.Data); err != nil {
		w.logger.Errorf("Error parsing Data field: %v. Raw JSON Data: %s", err, jsonData)
		return
	}

	change.Event = eventType

	if w.validLanguage(change) {
		changesChannel <- change
	}
}

func (w *wikipediaService) validLanguage(change dto.RecentChange) bool {
	if w.language == "" {
		return true
	}

	matches := re.FindStringSubmatch(change.Data.ParsedComment)
	return len(matches) > 1 && w.language == matches[1]
}

func (w *wikipediaService) doRecentChange() *http.Response {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(w.cfg.WikipediaConfig.StreamDataUrl)
	if err != nil {
		w.logger.Errorf("Error fetching stream data: %v", err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		w.logger.Errorf("Error fetching stream data: %d", resp.StatusCode)
		resp.Body.Close()
		return nil
	}

	return resp
}
