package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"

type IRemoveColumn interface {
    Execute(columnId int64)
}

type removeColumn struct {
    repository db.IColumnRepository
}

func NewRemoveColumn(repository db.IColumnRepository) IRemoveColumn {
    return &removeColumn{
        repository,
    }
}

func (rc *removeColumn) Execute(columnId int64) {
    rc.repository.Delete(columnId)
}