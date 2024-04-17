-- +goose Up

CREATE TABLE posts
(
    id           UUID PRIMARY KEY,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title        TEXT NOT NULL,
    url          TEXT NOT NULL UNIQUE,
    description  TEXT,
    published_at TIMESTAMP WITH TIME ZONE,
    feed_id      UUID NOT NULL REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;