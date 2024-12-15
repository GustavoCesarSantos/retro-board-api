CREATE TABLE IF NOT EXISTS teams (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);
