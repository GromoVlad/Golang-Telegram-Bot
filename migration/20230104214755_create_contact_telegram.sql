-- +goose Up
-- +goose StatementBegin
CREATE TABLE contact.telegram
(
    id            serial PRIMARY KEY,
    telegram_id   INTEGER UNIQUE          NOT NULL,
    telegram_name VARCHAR                 NOT NULL,
    is_subscriber boolean   DEFAULT false NOT NULL,
    created_at    TIMESTAMP DEFAULT now() NOT NULL,
    updated_at    TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contact.telegram;
-- +goose StatementEnd
