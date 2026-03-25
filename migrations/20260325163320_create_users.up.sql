CREATE TYPE user_role AS ENUM ('user', 'owner', 'admin');

CREATE TABLE IF NOT EXISTS users (
    uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    age INT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role user_role NOT NULL DEFAULT 'user',
    balance NUMERIC(10,2) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);