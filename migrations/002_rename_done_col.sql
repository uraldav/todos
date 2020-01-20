-- Write your migrate up statements here
-- +migrate Up
ALTER TABLE todo RENAME COLUMN done TO is_done;
---- create above / drop below ----
-- +migrate Down
ALTER TABLE todo RENAME COLUMN is_done TO done;
