CREATE TABLE IF NOT EXISTS polls (
    id bigserial PRIMARY KEY,
    poll_id BIGINT NOT NULL,
    text VARCHAR(50) NOT NULL,
    FOREIGN KEY (poll_id) REFERENCES polls (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);
