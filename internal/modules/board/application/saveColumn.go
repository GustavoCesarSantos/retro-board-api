package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type ISaveColumn interface {
    Execute(boardId int64, name string, color string, position int) (*domain.Column, error)
}

type saveColumn struct {
    repository db.IColumnRepository
}

func NewSaveColumn(repository db.IColumnRepository) ISaveColumn {
    return &saveColumn{
        repository,
    }
}

func (sc *saveColumn) Execute(boardId int64, name string, color string, position int) (*domain.Column, error) {
    column := domain.NewColumn(0, boardId, name, color, position)
    err := sc.repository.Save(column)
    return column, err
}
