CREATE TABLE IF NOT EXISTS board_columns (
    id bigserial PRIMARY KEY,
    board_id BIGINT NOT NULL,
    name text NOT NULL,
    color varchar(20) NOT NULL,
    position integer NOT NULL,
    FOREIGN KEY (board_id) REFERENCES boards (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);
