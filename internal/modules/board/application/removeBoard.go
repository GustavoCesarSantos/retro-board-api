package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"

type IRemoveBoard interface {
    Execute(boardId int64)
}

type removeBoard struct {
    repository db.IBoardRepository
}

func NewRemoveBoard(repository db.IBoardRepository) IRemoveBoard {
    return &removeBoard{
        repository,
    }
}

func (rb *removeBoard) Execute(boardId int64) {
    rb.repository.Delete(boardId)
}
