CREATE TABLE IF NOT EXISTS teamMembers (
    id bigserial PRIMARY KEY,
    teamId BIGINT NOT NULL,
    memberId BIGINT NOT NULL,
    role VARCHAR(50) NOT NULL,
    UNIQUE (teamId, memberId),
    FOREIGN KEY (teamId) REFERENCES teams (id) ON DELETE CASCADE,
    FOREIGN KEY (memberId) REFERENCES users (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);