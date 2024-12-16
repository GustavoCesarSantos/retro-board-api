package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type ISaveBoard interface {
    Execute(teamId int64, name string) error
}

type saveBoard struct {
    repository db.IBoardRepository
}

func NewSaveBoard(repository db.IBoardRepository) ISaveBoard {
    return &saveBoard{
        repository,
    }
}

func (sb *saveBoard) Execute(teamId int64, name string) error {
    board := domain.NewBoard(0, teamId, name)
    return sb.repository.Save(board)
}
