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

func (br *boardRepository) FindAllByTeamId(teamId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Board], error) {
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
            AND id < $2
        ORDER BY
            id DESC
        LIMIT $3;
    `
    args := []any{teamId, lastId, limit}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := br.DB.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    boards := []domain.Board{}
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
        boards = append(boards, board)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    var nextCursor int
	if len(boards) > 0 {
		nextCursor = int(boards[len(boards)-1].ID)
	}
    return &utils.ResultPaginated[domain.Board]{
        Items: boards,
        NextCursor: nextCursor,
    }, nil
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
            team_id,
            name,
            active
        )
        VALUES (
            $1,
            $2,
            true
        )
        RETURNING
            id,
            active,
            created_at
    `
    args := []any{
        board.TeamId,
        board.Name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return br.DB.QueryRowContext(ctx, query, args...).Scan(
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
