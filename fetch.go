package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func fetchTimeline(c config, lastID int) []byte {
	req, err := http.NewRequest(http.MethodGet, c.App.ServerURL+"/api/v1/timelines/home?"+url.Values{
		"limit":  []string{"40"},
		"min_id": []string{strconv.Itoa(lastID)},
	}.Encode(), nil)
	if err != nil {
		fatal("failed to fetch timeline: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.App.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fatal("failed to fetch home timeline: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fatal("failed to fetch home timeline: %v", resp.Status)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		fatal("failed to fetch home timeline: %v", err)

	}

	return buf.Bytes()
}

func parseTimeline(b []byte) []rawMessage {
	var messages []json.RawMessage
	if err := json.Unmarshal(b, &messages); err != nil {
		fatal("failed to parse home timeline: %v", err)
	}

	var out []rawMessage
	for _, msg := range messages {
		var parsedMessage struct {
			ID string `json:"id"`
		}
		if err := json.Unmarshal(msg, &parsedMessage); err != nil {
			fatal("failed to parse home timeline: %v", err)
		}
		id, err := strconv.ParseInt(parsedMessage.ID, 10, 64)
		if err != nil {
			fatal("failed to parse home timeline: %v", err)
		}
		out = append(out, rawMessage{
			id:      int(id),
			content: string(msg),
		})
	}
	return out
}

func fetch(c config, db *sql.DB) {
	for {
		lastID := lastSavedMessageID(db)

		timeline := fetchTimeline(c, lastID)
		messages := parseTimeline(timeline)
		for _, message := range messages {
			saveMessage(db, message)
		}
		if len(messages) == 0 {
			break
		}
	}
}
