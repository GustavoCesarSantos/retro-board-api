package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IGetNextColumnPosition interface {
    Execute(boardId int64) (int, error)
}

type getNextColumnPosition struct {
    repository db.IColumnRepository
}

func NewGetNextColumnPosition(repository db.IColumnRepository) IGetNextColumnPosition {
    return &getNextColumnPosition{
        repository,
    }
}

func (gnp *getNextColumnPosition) Execute(boardId int64) (int, error) {
    length, countErr := gnp.repository.CountColumnsByBoardId(boardId)
    if countErr != nil {
        return 0, countErr
    }
    return length + 1, nil
}
