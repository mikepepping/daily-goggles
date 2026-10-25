package store

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct{ *sql.DB }

func Open(path string) (*DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (name TEXT NOT NULL, state TEXT NOT NULL, done_at TEXT NOT NULL DEFAULT '');
		CREATE TABLE IF NOT EXISTS history (name TEXT NOT NULL, state TEXT NOT NULL, done_at TEXT NOT NULL DEFAULT '');
	`)
	if err != nil {
		db.Close()
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Tx(ctx context.Context, fn func(TaskRepo, HistoryRepo) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err = fn(NewTaskRepo(tx), NewHistoryRepo(tx)); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
