-- +goose Up

CREATE TABLE feed_follows
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    feed_id    UUID REFERENCES feeds (id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;