package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type UpdateColumnParams struct {
	Name *string
    Color *string
    Position *int
}

type IColumnRepository interface {
	CountColumnsByBoardId(boardId int64) (int, error)
	Delete(columnId int64) error
    FindAllByBoardId(boardId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Column], error)
    FindById(columnId int64) (*domain.Column, error)
	MoveOtherColumnsToLeftByColumnId(columnId int64, positionFrom int, positionTo int) error
    MoveOtherColumnsToRightByColumnId(columnId int64, positionFrom int, positionTo int) error
	Save(column *domain.Column) error
	Update(columnId int64, column UpdateColumnParams) error
}
