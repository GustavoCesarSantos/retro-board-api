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
           board_columns 
        WHERE
            board_id = $1
    `
    total := 0
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := cr.DB.QueryRowContext(ctx, query, boardId).Scan(
        &total,
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
            board_columns
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

func (cr *columnRepository) FindAllByBoardId(boardId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Column], error) {
    query := `
        SELECT
            id,
            name,
            color,
            position,
            created_at,
            updated_at
        FROM
            board_columns
        WHERE
            board_id = $1
            AND id < $2
        ORDER BY
            id DESC
        LIMIT $3;
    `
    args := []any{boardId, lastId, limit}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := cr.DB.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    columns := []domain.Column{}
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
        columns = append(columns, column)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    var nextCursor int
	if len(columns) > 0 {
		nextCursor = int(columns[len(columns)-1].ID)
	}
    return &utils.ResultPaginated[domain.Column]{
        Items: columns,
        NextCursor: nextCursor,
    }, nil
}

func (cr *columnRepository) FindById(columnId int64) (*domain.Column, error) {
    query := `
        SELECT
            id,
            name,
            color,
            position,
            created_at,
            updated_at
        FROM
            board_columns
        WHERE
            id = $1;
    `
	var column domain.Column
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := cr.DB.QueryRowContext(ctx, query, columnId).Scan(
		&column.ID,
        &column.Name,
        &column.Color,
        &column.Position,
		&column.CreatedAt,
		&column.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &column, nil
}

func (cr *columnRepository) MoveOtherColumnsToLeftByColumnId(columnId int64, positionFrom int, positionTo int) error {
    query := `
        UPDATE
           board_columns 
        SET
            position = position - 1,
            updated_at = NOW()
        WHERE
            id <> $1
            AND position > $2
            AND position <= $3;
    `
	args := []any{
        columnId,
        positionFrom,
        positionTo,
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

func (cr *columnRepository) MoveOtherColumnsToRightByColumnId(columnId int64, positionFrom int, positionTo int) error {
    query := `
        UPDATE
           board_columns 
        SET
            position = position + 1,
            updated_at = NOW()
        WHERE
            id <> $1
            AND position < $2
            AND position >= $3;
    `
	args := []any{
        columnId,
        positionFrom,
        positionTo,
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

func (cr *columnRepository) Save(column *domain.Column) error {
    query := `
        INSERT INTO board_columns (
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
    if column.Name == nil && column.Color == nil && column.Position == nil {
		return errors.New("NO COLUMN FIELD PROVIDED FOR UPDATING")
	}
    query := "UPDATE board_columns SET"
	var args []interface{}
	argPos := 1
    if column.Name != nil {
		query += " name = $" + strconv.Itoa(argPos) + ","
		args = append(args, *column.Name)
		argPos++
	}
	if column.Color != nil {
		query += " color = $" + strconv.Itoa(argPos) + ","
		args = append(args, *column.Color)
		argPos++
	}
    if column.Position != nil {
		query += " position = $" + strconv.Itoa(argPos) + ","
		args = append(args, *column.Position)
		argPos++
	}
    query += " updated_at = NOW()"
    query = strings.TrimSuffix(query, ",") + " WHERE id = $" + strconv.Itoa(argPos)
	args = append(args, columnId)
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
