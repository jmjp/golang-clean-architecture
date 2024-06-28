DROP SCHEMA  IF EXISTS public CASCADE;
CREATE SCHEMA public;

CREATE TABLE IF NOT EXISTS "users" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(20) UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    blocked BOOLEAN DEFAULT false,
    avatar TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create otps table if not exists
CREATE TABLE IF NOT EXISTS "otps" (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code VARCHAR(4) NOT NULL,
    valid_until TIMESTAMPTZ NOT NULL,
    UNIQUE (user_id, code)
);

CREATE TABLE IF NOT EXISTS "sessions" (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    jid UUID NOT NULL,
    ip INET NOT NULL,
    valid_until TIMESTAMPTZ NOT NULL,
    UNIQUE (user_id, jid)
);