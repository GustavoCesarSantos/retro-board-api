package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IUpdateBoard interface {
    Execute(boardId int64, board db.UpdateBoardParams) error
}

type updateBoard struct {
    repository db.IBoardRepository
}

func NewUpdateBoard(repository db.IBoardRepository) IUpdateBoard {
    return &updateBoard{
        repository,
    }
}

func (ub *updateBoard) Execute(boardId int64, board db.UpdateBoardParams) error {
    return ub.repository.Update(boardId, board)
}
