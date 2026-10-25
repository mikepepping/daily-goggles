package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/mikepepping/daily-goggles/tasks"
)

type taskDB interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

type TaskRepo struct{ db taskDB }

func NewTaskRepo(db taskDB) TaskRepo { return TaskRepo{db: db} }

func (r TaskRepo) List(ctx context.Context) ([]tasks.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT name, state, done_at FROM tasks ORDER BY rowid")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

func (r TaskRepo) Get(ctx context.Context, index int) (tasks.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT name, state, done_at FROM tasks ORDER BY rowid LIMIT 1 OFFSET ?", index)
	if err != nil {
		return tasks.Task{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return tasks.Task{}, sql.ErrNoRows
	}
	return scanTask(rows)
}

func (r TaskRepo) Create(ctx context.Context, task tasks.Task) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tasks (name, state, done_at) VALUES (?, ?, ?)", task.Name, task.State, date(task.DoneAt))
	return err
}

func (r TaskRepo) Update(ctx context.Context, index int, task tasks.Task) error {
	_, err := r.db.ExecContext(ctx, "UPDATE tasks SET name = ?, state = ?, done_at = ? WHERE rowid = (SELECT rowid FROM tasks ORDER BY rowid LIMIT 1 OFFSET ?)", task.Name, task.State, date(task.DoneAt), index)
	return err
}

func (r TaskRepo) Delete(ctx context.Context, index int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tasks WHERE rowid = (SELECT rowid FROM tasks ORDER BY rowid LIMIT 1 OFFSET ?)", index)
	return err
}

func scanRows(rows *sql.Rows) ([]tasks.Task, error) {
	var result []tasks.Task
	for rows.Next() {
		task, err := scanTask(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, task)
	}
	return result, rows.Err()
}

type scanner interface{ Scan(...any) error }

func scanTask(s scanner) (tasks.Task, error) {
	var task tasks.Task
	var doneAt string
	if err := s.Scan(&task.Name, &task.State, &doneAt); err != nil {
		return tasks.Task{}, err
	}
	if doneAt != "" {
		task.DoneAt, _ = time.Parse(time.RFC3339Nano, doneAt)
	}
	return task, nil
}

func date(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339Nano)
}
