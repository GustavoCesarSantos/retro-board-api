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

type pollRepository struct {
	DB *sql.DB
}

func NewPollRepository(db *sql.DB) db.IPollRepository {
	return &pollRepository{
		DB: db,
	}
}

func (pr *pollRepository) FindAllByTeamId(teamId int64) ([]*domain.Poll, error) {
	query := `
        SELECT
            id,
            team_id,
			name,
            created_at,
            updated_at
        FROM
            polls
        WHERE
            team_id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := pr.DB.QueryContext(ctx, query, teamId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    polls := []*domain.Poll{}
    for rows.Next() {
        var poll domain.Poll
        err := rows.Scan(
            &poll.ID,
            &poll.TeamId,
            &poll.Name,
            &poll.CreatedAt,
            &poll.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        polls = append(polls, &poll)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return polls, nil
}

func (pr *pollRepository) FindById(pollId int64) (*domain.Poll, error) {
	query := `
        SELECT
            id,
            team_id,
			name,
            created_at,
            updated_at
        FROM
            polls
        WHERE
            id = $1;
    `
	var poll domain.Poll
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := pr.DB.QueryRowContext(ctx, query, pollId).Scan(
		&poll.ID,
        &poll.Name,
		&poll.CreatedAt,
		&poll.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &poll, nil
}

func (pr *pollRepository) Save(poll *domain.Poll) error {
	query := `
        INSERT INTO polls (
            team_id,
			name
        )
        VALUES (
            $1,
			$2
        )
        RETURNING
            id,
            created_at
    `
	args := []any{poll.TeamId, poll.Name}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return pr.DB.QueryRowContext(ctx, query, args...).Scan(
        &poll.ID,
        &poll.CreatedAt,
    )
}
