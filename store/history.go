package store

import (
	"context"
	"database/sql"
	"github.com/mikepepping/daily-goggles/tasks"
)

type HistoryRepo struct{ db taskDB }

func NewHistoryRepo(db taskDB) HistoryRepo { return HistoryRepo{db: db} }

func (r HistoryRepo) List(ctx context.Context) ([]tasks.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT name, state, done_at FROM history ORDER BY rowid")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

func (r HistoryRepo) Get(ctx context.Context, index int) (tasks.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT name, state, done_at FROM history ORDER BY rowid LIMIT 1 OFFSET ?", index)
	if err != nil {
		return tasks.Task{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return tasks.Task{}, sql.ErrNoRows
	}
	return scanTask(rows)
}

func (r HistoryRepo) Create(ctx context.Context, task tasks.Task) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO history (name, state, done_at) VALUES (?, ?, ?)", task.Name, task.State, date(task.DoneAt))
	return err
}

func (r HistoryRepo) Update(ctx context.Context, index int, task tasks.Task) error {
	_, err := r.db.ExecContext(ctx, "UPDATE history SET name = ?, state = ?, done_at = ? WHERE rowid = (SELECT rowid FROM history ORDER BY rowid LIMIT 1 OFFSET ?)", task.Name, task.State, date(task.DoneAt), index)
	return err
}

func (r HistoryRepo) Delete(ctx context.Context, index int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM history WHERE rowid = (SELECT rowid FROM history ORDER BY rowid LIMIT 1 OFFSET ?)", index)
	return err
}
