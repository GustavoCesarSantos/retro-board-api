package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type voteRepository struct {
	DB *sql.DB
}

func NewVoteRepository(db *sql.DB) db.IVoteRepository {
	return &voteRepository{
		DB: db,
	}
}

func (vr *voteRepository) CountByOptionId(optionId int64) (int, error) {
	query := `
        SELECT
			COUNT(*) AS count
        FROM
            poll_votes
        WHERE
            option_id = $1;
    `
	var count int
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := vr.DB.QueryRowContext(ctx, query, optionId).Scan(
		count,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, utils.ErrRecordNotFound
		default:
			return 0, err
		}
	}
	return count, nil
}

func (vr *voteRepository) Save(vote *domain.Vote) error {
	query := `
        INSERT INTO poll_votes (
            member_id,
			option_id
        )
        VALUES (
            $1,
			$2
        )
        RETURNING
            id,
            created_at
    `
	args := []any{vote.MemberId, vote.OptionId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return vr.DB.QueryRowContext(ctx, query, args...).Scan(
        &vote.ID,
        &vote.CreatedAt,
    )
}
