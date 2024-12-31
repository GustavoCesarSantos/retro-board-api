package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type optionRepository struct {
	DB *sql.DB
}

func NewOptionRepository(db *sql.DB) db.IOptionRepository {
	return &optionRepository{
		DB: db,
	}
}

func (or *optionRepository) Delete(optionId int64) error {
	query := `
        DELETE FROM 
            poll_options
        WHERE 
            id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := or.DB.ExecContext(ctx, query, optionId)
    if err != nil {
        return err
    }
    rowsAffected, rowsAffectedErr := result.RowsAffected()
    if rowsAffectedErr != nil {
        return rowsAffectedErr
    }
    if rowsAffected == 0 {
        return utils.ErrRecordNotFound
    }
    return nil
}

func (or *optionRepository) FindAllByPollId(pollId int64) ([]*domain.Option, error) {
	query := `
        SELECT
            id,
            poll_id,
			text,
            created_at,
            updated_at
        FROM
            poll_options
        WHERE
            poll_id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := or.DB.QueryContext(ctx, query, pollId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    options := []*domain.Option{}
    for rows.Next() {
        var option domain.Option
        err := rows.Scan(
            &option.ID,
            &option.PollId,
            &option.Text,
            &option.CreatedAt,
            &option.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        options = append(options, &option)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return options, nil
}

func (or *optionRepository) Save(option *domain.Option) error {
	query := `
        INSERT INTO poll_options (
            poll_id,
			text
        )
        VALUES (
            $1,
			$2
        )
        RETURNING
            id,
            created_at
    `
	args := []any{option.PollId, option.Text}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return or.DB.QueryRowContext(ctx, query, args...).Scan(
        &option.ID,
        &option.CreatedAt,
    )
}
