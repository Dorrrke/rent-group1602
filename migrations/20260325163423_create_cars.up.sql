CREATE TABLE IF NOT EXISTS cars (
    cid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    color TEXT NOT NULL,
    year INT NOT NULL,
    number TEXT UNIQUE NOT NULL,
    price_per_day NUMERIC(10,2) NOT NULL,
    available BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);