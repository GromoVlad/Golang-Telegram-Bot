-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA contact;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA contact;
-- +goose StatementEnd
