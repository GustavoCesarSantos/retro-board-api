package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"

type IRemoveBoard interface {
    Execute(boardId int64) error
}

type removeBoard struct {
    repository db.IBoardRepository
}

func NewRemoveBoard(repository db.IBoardRepository) IRemoveBoard {
    return &removeBoard{
        repository,
    }
}

func (rb *removeBoard) Execute(boardId int64) error {
    return rb.repository.Delete(boardId)
}
