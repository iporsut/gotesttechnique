-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos;
-- +goose StatementEnd
