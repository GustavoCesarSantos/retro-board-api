CREATE TABLE IF NOT EXISTS board_cards (
    id bigserial PRIMARY KEY,
    column_id BIGINT NOT NULL,
    member_id BIGINT NOT NULL,
    text text NOT NULL,
    FOREIGN KEY (column_id) REFERENCES board_columns (id) ON DELETE CASCADE,
    FOREIGN KEY (member_id) REFERENCES users (id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);
