CREATE TABLE IF NOT EXISTS team_members (
    id bigserial PRIMARY KEY,
    team_id BIGINT NOT NULL,
    member_id BIGINT NOT NULL,
    role VARCHAR(50) NOT NULL,
    UNIQUE (team_id, member_id),
    FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    FOREIGN KEY (member_id) REFERENCES users (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);