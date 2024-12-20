CREATE TABLE IF NOT EXISTS boards (
    id bigserial PRIMARY KEY,
    team_id BIGINT NOT NULL,
    name text NOT NULL,
    active boolean NOT NULL,
    FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);
