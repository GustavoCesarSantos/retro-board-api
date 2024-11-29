package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IGetNextColumnPosition interface {
    Execute(boardId int64) int
}

type getNextColumnPosition struct {
    repository db.IColumnRepository
}

func NewGetNextColumnPosition(repository db.IColumnRepository) IGetNextColumnPosition {
    return &getNextColumnPosition{
        repository,
    }
}

func (gnp *getNextColumnPosition) Execute(boardId int64) int {
    length := gnp.repository.CountColumnsByBoardId(boardId)
    return length + 1
}
