-- name: GetTodo :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todo
ORDER BY done;

-- name: CreateTodo :one
INSERT INTO todo (text)
VALUES ($1) RETURNING *;

-- name: UpdateTodoDone :one
UPDATE todo
SET done = $2
WHERE id = $1
RETURNING *;
