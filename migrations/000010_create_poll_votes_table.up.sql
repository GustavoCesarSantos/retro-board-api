CREATE TABLE IF NOT EXISTS poll_votes (
    id bigserial PRIMARY KEY,
    member_id BIGINT NOT NULL,
    option_id BIGINT NOT NULL,
    FOREIGN KEY (member_id) REFERENCES team_members (id) ON DELETE CASCADE,
    FOREIGN KEY (option_id) REFERENCES poll_options (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
