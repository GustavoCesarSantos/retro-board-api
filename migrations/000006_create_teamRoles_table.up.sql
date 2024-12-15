CREATE TABLE IF NOT EXISTS teamRoles (
    id bigserial PRIMARY KEY,
    role VARCHAR(50) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);