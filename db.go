package main

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/dottedmag/sqv"
	_ "github.com/glebarez/go-sqlite"
)

const masemaAppID = 0x6d61736d

var schema = []string{
	`
CREATE TABLE messages (
  id      INTEGER PRIMARY KEY,
  content TEXT,
  sent    INTEGER
)
`,
}

func defaultDBFile() string {
	return xdgDataDir() + "/masema/messages.db"
}

func openDB(dbFileName string) (*sql.DB, error) {
	os.MkdirAll(filepath.Dir(dbFileName), 0o755)

	db, err := sql.Open("sqlite", dbFileName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}
	if _, err := db.Exec("PRAGMA journal_mode = WAL"); err != nil {
		return nil, err
	}

	if err := sqv.Apply(context.Background(), db, masemaAppID, schema); err != nil {
		return nil, err
	}

	return db, nil
}

func lastSavedMessageID(db *sql.DB) int {
	var id int
	if err := db.QueryRow("SELECT id FROM messages ORDER BY id DESC LIMIT 1").Scan(&id); err != nil {
		if err != sql.ErrNoRows {
			fatal("failed to get last message ID: %v", err)
		}
		return 0
	}
	return id
}

func saveMessage(db *sql.DB, msg rawMessage) {
	if _, err := db.Exec("INSERT INTO messages (id, content) VALUES (?, ?)", msg.id, msg.content); err != nil {
		fatal("failed to save message: %v", err)
	}
}

func unsentMessages(db *sql.DB) []rawMessage {
	rows, err := db.Query("SELECT id, content FROM messages WHERE sent IS NULL")
	if err != nil {
		fatal("failed to fetch unsent messages: %s", err)
	}
	defer rows.Close()

	var msgs []rawMessage
	for rows.Next() {
		var msg rawMessage
		if err := rows.Scan(&msg.id, &msg.content); err != nil {
			fatal("failed to fetch unsent messages: %s", err)
		}
		msgs = append(msgs, msg)
	}
	return msgs
}

func markMessageAsSent(db *sql.DB, id int) {
	if _, err := db.Exec("UPDATE messages SET sent = 1 WHERE id = ?", id); err != nil {
		fatal("failed to mark message %d as sent: %s", id, err)
	}
}
