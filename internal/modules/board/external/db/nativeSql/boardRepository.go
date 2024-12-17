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

type boardRepository struct {
    DB *sql.DB
}

func NewBoardRepository(db *sql.DB) db.IBoardRepository {
	return &boardRepository{
        DB: db,
	}
}

func (br *boardRepository) Delete(boardId int64) error { 
    query := `
        DELETE FROM
            boards
        WHERE
            id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := br.DB.ExecContext(ctx, query, boardId)
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

func (br *boardRepository) FindAllByTeamId(teamId int64) ([]*domain.Board, error) {
    query := `
        SELECT
            id,
            name,
            active,
            created_at,
            updated_at
        FROM
            boards
        WHERE
            team_id = $1
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := br.DB.QueryContext(ctx, query, teamId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    boards := []*domain.Board{}
    for rows.Next() {
        var board domain.Board
        err := rows.Scan(
            &board.ID,
            &board.Name,
            &board.Active,
            &board.CreatedAt,
            &board.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        boards = append(boards, &board)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return boards, nil
}

func (br *boardRepository) FindById(boardId int64) (*domain.Board, error) {
    query := `
        SELECT
            id,
            name,
            active,
            created_at,
            updated_at
        FROM
            boards 
        WHERE
            id = $1
    `
	var board domain.Board
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := br.DB.QueryRowContext(ctx, query, boardId).Scan(
		&board.ID,
        &board.Name,
        &board.Active,
		&board.CreatedAt,
		&board.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &board, nil
}

func (br *boardRepository) Save(board *domain.Board) error {
    query := `
        INSERT INTO boards (
            name,
            active
        )
        VALUES (
            $1,
            true
        )
        RETURNING
            id,
            active,
            created_at
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return br.DB.QueryRowContext(ctx, query, board.Name).Scan(
        &board.ID,
        &board.Active,
        &board.CreatedAt,
    )
}

func (br *boardRepository) Update(boardId int64, board db.UpdateBoardParams) error {
    query := `
        UPDATE
            boards
        SET
            name = $1,
            active = $2
        WHERE
            id = $3;
    `
	args := []any{
        board.Name,
        board.Active,
        boardId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := br.DB.ExecContext(ctx, query, args...)
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
