package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IFindBoard interface {
    Execute(boardId int64) *domain.Board
}

type findBoard struct {
    repository db.IBoardRepository
}

func NewFindBoard(repository db.IBoardRepository) IFindBoard {
    return &findBoard{
        repository,
    }
}

func (fb *findBoard) Execute(boardId int64) *domain.Board {
    return fb.repository.FindById(boardId)
}