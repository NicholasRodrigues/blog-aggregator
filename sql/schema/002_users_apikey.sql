-- +goose Up
ALTER TABLE users ADD COLUMN apikey VARCHAR(64) NOT NULL UNIQUE
DEFAULT encode(sha256(random()::text::bytea), 'hex');