package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type UpdateBoardParams struct {
	Name *string
	Active *bool
}

type IBoardRepository interface {
	Delete(boardId int64) error
    FindAllByTeamId(teamId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Board], error)
    FindById(boardId int64) (*domain.Board, error)
	Save(board *domain.Board) error
	Update(boardId int64, board UpdateBoardParams) error
}
