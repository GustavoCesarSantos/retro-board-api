package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IUpdateColumn interface {
    Execute(columnId int64, column db.UpdateColumnParams)
}

type updateColumn struct {
    repository db.IColumnRepository
}

func NewUpdateColumn(repository db.IColumnRepository) IUpdateColumn {
    return &updateColumn{
        repository,
    }
}

func (uc *updateColumn) Execute(columnId int64, column db.UpdateColumnParams) {
    uc.repository.Update(columnId, column)
}
