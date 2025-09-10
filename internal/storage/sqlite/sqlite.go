package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/cmczk/shortest-url/internal/storage"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: cannot connect to database: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: database does not response: %w", op, err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		alias TEXT NOT NULL UNIQUE,
		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);`)
	if err != nil {
		return nil, fmt.Errorf("%s: cannot create tables urls: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(newURL, alias string) (int64, error) {
	const op = "storage.sqlite.Save"

	stmt, err := s.db.Prepare(`INSERT INTO urls (url, alias) VALUES (?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s: cannot prepare statement saving new url: %w", op, err)
	}

	res, err := stmt.Exec(newURL, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, storage.ErrURLExists
		}

		return 0, fmt.Errorf("%s: cannot save new url: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: cannot get id of new url: %w", op, err)
	}

	return id, nil
}
