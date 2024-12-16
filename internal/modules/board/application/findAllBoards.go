package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IFindAllBoards interface {
    Execute(teamId int64) ([]*domain.Board, error)
}

type findAllBoards struct {
    repository db.IBoardRepository
}

func NewFindAllBoards(repository db.IBoardRepository) IFindAllBoards {
    return &findAllBoards{
        repository,
    }
}

func (fab *findAllBoards) Execute(teamId int64) ([]*domain.Board, error) {
    return fab.repository.FindAllByTeamId(teamId)
}
