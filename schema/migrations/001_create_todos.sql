-- Write your migrate up statements here
-- +migrate Up
CREATE SEQUENCE todo_id_seq INCREMENT BY 1 MINVALUE 1 START 1;

CREATE TABLE todo (
  id    INT       NOT NULL DEFAULT nextval('todo_id_seq'::regclass),
  text  text      NOT NULL,
  done  BOOLEAN   NOT NULL DEFAULT FALSE,
  PRIMARY KEY(id)
);
---- create above / drop below ----
-- +migrate Down
DROP TABLE todo;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
