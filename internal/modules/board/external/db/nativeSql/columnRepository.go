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

type columnRepository struct {
    DB *sql.DB
}

func NewColumnRepository(db *sql.DB) db.IColumnRepository {
	return &columnRepository{
        DB: db,
	}
}

func (cr *columnRepository) CountColumnsByBoardId(boardId int64) (int, error) {
    query := `
        SELECT
            COUNT(*) AS total
        FROM
           columns 
        WHERE
            board_id = $1
    `
    total := 0
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := cr.DB.QueryRowContext(ctx, query, boardId).Scan(
        total,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, utils.ErrRecordNotFound
		default:
			return 0, err
		}
	}
	return total, nil
}

func (cr *columnRepository) Delete(columnId int64) error {
    query := `
        DELETE FROM
            columns
        WHERE
            id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := cr.DB.ExecContext(ctx, query, columnId)
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

func (cr *columnRepository) FindAllByBoardId(boardId int64) ([]*domain.Column, error) {
    query := `
        SELECT
            id,
            name,
            color,
            position,
            created_at,
            updated_at
        FROM
            columns
        WHERE
            board_id = $1
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := cr.DB.QueryContext(ctx, query, boardId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    columns := []*domain.Column{}
    for rows.Next() {
        var column domain.Column
        err := rows.Scan(
            &column.ID,
            &column.Name,
            &column.Color,
            &column.Position,
            &column.CreatedAt,
            &column.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        columns = append(columns, &column)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return columns, nil
}

func (cr *columnRepository) Save(column *domain.Column) error {
    query := `
        INSERT INTO columns (
            board_id,
            name,
            color,
            position
        )
        VALUES (
            $1,
            $2,
            $3,
            $4
        )
        RETURNING
            id,
            created_at
    `
    args := []any{column.BoardId, column.Name, column.Color, column.Position}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return cr.DB.QueryRowContext(ctx, query, args...).Scan(
        &column.ID,
        &column.CreatedAt,
    )
}

func (cr *columnRepository) Update(columnId int64, column db.UpdateColumnParams) error {
    query := `
        UPDATE
           cards 
        SET
            name = $1,
            color = $2
        WHERE
            id = $3;
    `
	args := []any{
        column.Name,
        column.Color,
        columnId,
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