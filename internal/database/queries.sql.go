// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package database

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todo (text) VALUES ($1) RETURNING id, text, done
`

func (q *Queries) CreateTodo(ctx context.Context, text string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, text)
	var i Todo
	err := row.Scan(&i.ID, &i.Text, &i.Done)
	return i, err
}

const getTodo = `-- name: GetTodo :one
SELECT id, text, done FROM todo
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Text, &i.Done)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, text, done FROM todo
ORDER BY done
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Text, &i.Done); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodoDone = `-- name: UpdateTodoDone :one
UPDATE todo
SET done = $2
WHERE id = $1
RETURNING id, text, done
`

type UpdateTodoDoneParams struct {
	ID   int32 `json:"id"`
	Done bool  `json:"done"`
}

func (q *Queries) UpdateTodoDone(ctx context.Context, arg UpdateTodoDoneParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodoDone, arg.ID, arg.Done)
	var i Todo
	err := row.Scan(&i.ID, &i.Text, &i.Done)
	return i, err
}