-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    age SMALLINT NOT NULL,
    nationality VARCHAR(255) NOT NULL,
    sex VARCHAR(7) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS persons;
-- +goose StatementEnd
