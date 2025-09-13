-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS temperatures (
    id SERIAL PRIMARY KEY,
    sensor_id TEXT NOT NULL,
    temperature_celsius FLOAT NOT NULL,
    reading_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS temperatures;
-- +goose StatementEnd
