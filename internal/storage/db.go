package storage

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	dbPath := os.Getenv("PEA_DB_PATH")
	if dbPath == "" {
		appData := os.Getenv("APPDATA")
		if appData != "" {
			dbPath = appData + "/pea/pea.db"
		} else {
			dbPath = "./data/pea.db"
		}
	}

	if err := os.MkdirAll(dbPath[:len(dbPath)-len("pea.db")], 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS conversations (
			id TEXT PRIMARY KEY,
			title TEXT,
			created_at TEXT,
			updated_at TEXT
		);

		CREATE TABLE IF NOT EXISTS messages (
			id TEXT PRIMARY KEY,
			conversation_id TEXT,
			role TEXT,
			content TEXT,
			attachments JSON,
			model TEXT,
			model_config JSON,
			prompt_version TEXT,
			metadata JSON,
			created_at TEXT
		);

		CREATE TABLE IF NOT EXISTS prompts (
			id TEXT PRIMARY KEY,
			version TEXT,
			author TEXT,
			intent TEXT,
			description TEXT,
			prompt_text TEXT,
			settings JSON,
			examples JSON,
			created_at TEXT
		);

		CREATE TABLE IF NOT EXISTS runs (
			id TEXT PRIMARY KEY,
			prompt_id TEXT,
			model TEXT,
			settings JSON,
			status TEXT,
			metadata JSON,
			ephemeral_token TEXT,
			created_at TEXT
		);

		CREATE TABLE IF NOT EXISTS evaluations (
			id TEXT PRIMARY KEY,
			run_id TEXT,
			metrics JSON,
			created_at TEXT
		);

		CREATE TABLE IF NOT EXISTS attachments (
			id TEXT PRIMARY KEY,
			filename TEXT,
			path TEXT,
			mimetype TEXT,
			text_extract TEXT,
			created_at TEXT
		);

		CREATE TABLE IF NOT EXISTS audit_logs (
			id TEXT PRIMARY KEY,
			action TEXT,
			actor TEXT,
			details JSON,
			created_at TEXT
		);

		CREATE VIRTUAL TABLE IF NOT EXISTS prompts_fts USING fts5(prompt_text, content='prompts', content_rowid='rowid');
		CREATE VIRTUAL TABLE IF NOT EXISTS messages_fts USING fts5(content, content='messages', content_rowid='rowid');
	`); err != nil {
		return nil, err
	}

	return db, nil
}
