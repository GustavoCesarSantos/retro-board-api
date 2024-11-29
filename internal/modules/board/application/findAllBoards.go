package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IFindAllBoards interface {
    Execute(teamId int64) []*domain.Board
}

type findAllBoards struct {
    repository db.IBoardRepository
}

func NewFindAllBoards(repository db.IBoardRepository) IFindAllBoards {
    return &findAllBoards{
        repository,
    }
}

func (fab *findAllBoards) Execute(teamId int64) []*domain.Board {
    return fab.repository.FindAllByTeamId(teamId)
}
