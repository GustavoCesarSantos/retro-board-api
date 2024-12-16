package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IFindAllColumns interface {
    Execute(boardId int64) ([]*domain.Column, error)
}

type findAllColumns struct {
    repository db.IColumnRepository
}

func NewFindAllColumns(repository db.IColumnRepository) IFindAllColumns {
    return &findAllColumns{
        repository,
    }
}

func (fac *findAllColumns) Execute(boardId int64) ([]*domain.Column, error) {
    return fac.repository.FindAllByBoardId(boardId)
}
