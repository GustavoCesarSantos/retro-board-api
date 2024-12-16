package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IUpdateColumn interface {
    Execute(columnId int64, column db.UpdateColumnParams) error
}

type updateColumn struct {
    repository db.IColumnRepository
}

func NewUpdateColumn(repository db.IColumnRepository) IUpdateColumn {
    return &updateColumn{
        repository,
    }
}

func (uc *updateColumn) Execute(columnId int64, column db.UpdateColumnParams) error {
    return uc.repository.Update(columnId, column)
}
