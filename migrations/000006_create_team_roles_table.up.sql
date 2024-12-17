CREATE TABLE IF NOT EXISTS team_roles (
    id bigserial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);