CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE auth.auth_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    hashed_password TEXT NOT NULL
);