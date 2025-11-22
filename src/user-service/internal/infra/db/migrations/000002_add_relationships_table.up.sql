CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS community.relationships (
    follower_id UUID NOT NULL REFERENCES community.users(id) ON DELETE CASCADE,
    followee_id UUID NOT NULL REFERENCES community.users(id) ON DELETE CASCADE,
    are_friends BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (follower_id, followee_id),
    CONSTRAINT follower_no_self CHECK (follower_id <> followee_id)
);