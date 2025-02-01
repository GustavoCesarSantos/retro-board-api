package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type IFindAllColumns interface {
    Execute(boardId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Column], error)
}

type findAllColumns struct {
    repository db.IColumnRepository
}

func NewFindAllColumns(repository db.IColumnRepository) IFindAllColumns {
    return &findAllColumns{
        repository,
    }
}

func (fac *findAllColumns) Execute(boardId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Column], error) {
    return fac.repository.FindAllByBoardId(boardId, limit, lastId)
}
