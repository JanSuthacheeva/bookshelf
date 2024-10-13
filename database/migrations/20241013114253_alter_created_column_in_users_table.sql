-- +goose Up
-- +goose StatementBegin
ALTER TABLE users CHANGE created created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users CHANGE created_at created DATETIME NOT NULL;
-- +goose StatementEnd
