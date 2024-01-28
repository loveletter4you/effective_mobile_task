-- +goose Up
-- +goose StatementBegin
ALTER TABLE persons RENAME COLUMN sex TO gender;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE persons RENAME COLUMN gender TO sex;
-- +goose StatementEnd
