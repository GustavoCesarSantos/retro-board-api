package db

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
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
            board_cards
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

func (cr *cardRepository) FindAllByColumnId(columnId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Card], error) {
    query := `
        SELECT
            id,
            column_id,
            member_id,
            text,
            created_at,
            updated_at
        FROM
           board_cards 
        WHERE
            column_id = $1
            AND id < $2
        ORDER BY
            id DESC
        LIMIT $3;
    `
    args := []any{columnId, lastId, limit}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := cr.DB.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    cards := []domain.Card{}
    for rows.Next() {
        var card domain.Card
        err := rows.Scan(
            &card.ID,
            &card.ColumnId,
            &card.MemberId,
            &card.Text,
            &card.CreatedAt,
            &card.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    var nextCursor int
	if len(cards) > 0 {
		nextCursor = int(cards[len(cards)-1].ID)
	}
    return &utils.ResultPaginated[domain.Card]{
        Items: cards,
        NextCursor: nextCursor,
    }, nil
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
           board_cards 
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
        INSERT INTO board_cards (
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
    if card.Text == nil && card.ColumnId == nil {
		return errors.New("NO CARD FIELD PROVIDED FOR UPDATING")
	}
    query := "UPDATE board_cards SET"
	var args []interface{}
	argPos := 1
    if card.Text != nil {
		query += " text = $" + strconv.Itoa(argPos) + ","
		args = append(args, *card.Text)
		argPos++
	}
	if card.ColumnId != nil {
		query += " column_id = $" + strconv.Itoa(argPos) + ","
		args = append(args, *card.ColumnId)
		argPos++
	}
    query += " updated_at = NOW()"
    query = strings.TrimSuffix(query, ",") + " WHERE id = $" + strconv.Itoa(argPos)
	args = append(args, cardId)
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
