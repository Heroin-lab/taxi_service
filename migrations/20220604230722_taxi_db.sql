-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA taxi_db;

CREATE TABLE IF NOT EXISTS taxi_db.users(
    user_id INT PRIMARY KEY,
    login CHAR(255),
    password_hash CHAR(255));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE taxi_db.users;

DROP SCHEMA taxi_db CASCADE;
-- +goose StatementEnd
