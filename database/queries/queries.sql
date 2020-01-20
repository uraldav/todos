-- name: ListTodos :many
SELECT id, text, is_done FROM todo
ORDER BY id;

-- name: CreateTodo :one
INSERT INTO todo (text)
VALUES ($1) RETURNING *;

-- name: UpdateTodoDone :one
UPDATE todo
SET is_done = $2
WHERE id = $1
RETURNING id, text, is_done;

-- name: DeleteTodo :one
DELETE FROM todo
WHERE id = $1
RETURNING id, text, is_done;
