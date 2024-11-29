package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type ISaveBoard interface {
    Execute(teamId int64, name string)
}

type saveBoard struct {
    repository db.IBoardRepository
}

func NewSaveBoard(repository db.IBoardRepository) ISaveBoard {
    return &saveBoard{
        repository,
    }
}

func (sb *saveBoard) Execute(teamId int64, name string) {
    board := domain.NewBoard(0, teamId, name)
    sb.repository.Save(*board)
}
