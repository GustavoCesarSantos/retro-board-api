package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"

type UpdateBoardParams struct {
	Name *string
	Active *bool
}

type IBoardRepository interface {
	Delete(boardId int64) error
    FindAllByTeamId(teamId int64) ([]*domain.Board, error)
    FindById(boardId int64) (*domain.Board, error)
	Save(board *domain.Board) error
	Update(boardId int64, board UpdateBoardParams) error
}
