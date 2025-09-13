CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    post_karma INT DEFAULT 0,
    comment_karma INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT now()
);