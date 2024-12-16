package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)


type cardRepository struct {
    DB *sql.DB
}

func NewCardRepository(db *sql.DB) db.ICardRepository {
	return &cardRepository{
        DB: db,
	}
}

func (cr *cardRepository) Delete(cardId int64) error {
    query := `
        DELETE FROM
            cards
        WHERE
            id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := cr.DB.ExecContext(ctx, query, cardId)
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

func (cr *cardRepository) FindAllByColumnId(columnId int64) ([]*domain.Card, error) {
    query := `
        SELECT
            id,
            member_id,
            text,
            created_at,
            updated_at
        FROM
           cards 
        WHERE
            column_id = $1
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := cr.DB.QueryContext(ctx, query, columnId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    cards := []*domain.Card{}
    for rows.Next() {
        var card domain.Card
        err := rows.Scan(
            &card.ID,
            &card.MemberId,
            &card.Text,
            &card.CreatedAt,
            &card.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        cards = append(cards, &card)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return cards, nil
}

func (cr *cardRepository) FindById(cardId int64) (*domain.Card, error) {
    query := `
        SELECT
            id,
            column_id,
            member_id,
            text,
            created_at,
            updated_at
        FROM
           cards 
        WHERE
            id = $1
    `
	var card domain.Card
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := cr.DB.QueryRowContext(ctx, query, cardId).Scan(
		&card.ID,
        &card.ColumnId,
        &card.MemberId,
        &card.Text,
		&card.CreatedAt,
		&card.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &card, nil
}

func (cr *cardRepository) Save(card *domain.Card) error {
    query := `
        INSERT INTO cards (
            column_id,
            member_id,
            text
        )
        VALUES (
            $1,
            $2,
            $3
        )
        RETURNING
            id,
            created_at
    `
    args := []any{card.ColumnId, card.MemberId, card.Text}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return cr.DB.QueryRowContext(ctx, query, args...).Scan(
        &card.ID,
        &card.CreatedAt,
    )
}

func (cr *cardRepository) Update(cardId int64, card db.UpdateCardParams) error {
    query := `
        UPDATE
            cards 
        SET
            text = $1,
            column_id = $2
        WHERE
            id = $3;
    `
	args := []any{
        card.Text,
        card.ColumnId,
        cardId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := cr.DB.ExecContext(ctx, query, args...)
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
