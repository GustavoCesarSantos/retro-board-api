package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IFindColumn interface {
    Execute(columnId int64) (*domain.Column, error)
}

type findColumn struct {
    repository db.IColumnRepository
}

func NewFindColumn(repository db.IColumnRepository) IFindColumn {
    return &findColumn{
        repository,
    }
}

func (fc *findColumn) Execute(columnId int64) (*domain.Column, error) {
    return fc.repository.FindById(columnId)
}