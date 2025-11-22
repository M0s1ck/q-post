CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE auth.refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token_hash TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES auth.auth_users(id) ON DELETE CASCADE,
    expires_at TIMESTAMP NOT NULL
)