package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type ISaveBoard interface {
    Execute(teamId int64, name string) (*domain.Board, error)
}

type saveBoard struct {
    repository db.IBoardRepository
}

func NewSaveBoard(repository db.IBoardRepository) ISaveBoard {
    return &saveBoard{
        repository,
    }
}

func (sb *saveBoard) Execute(teamId int64, name string) (*domain.Board, error) {
    board := domain.NewBoard(0, teamId, name)
    err := sb.repository.Save(board)
    return board, err
}
