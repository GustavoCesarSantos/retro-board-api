package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type UpdateColumnParams struct {
	Name *string
    Color *string
}

type IColumnRepository interface {
	CountColumnsByBoardId(boardId int64) (int, error)
	Delete(columnId int64) error
    FindAllByBoardId(boardId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Column], error)
	Save(column *domain.Column) error
	Update(columnId int64, column UpdateColumnParams) error
}
