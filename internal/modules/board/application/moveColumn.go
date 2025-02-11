package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IMoveColumn interface {
    Execute(columnId int64, positionTo int) error
}

type moveColumn struct {
    repository db.IColumnRepository
}

func NewMoveColumn(repository db.IColumnRepository) IMoveColumn {
    return &moveColumn{
        repository,
    }
}

func (mc *moveColumn) Execute(columnId int64, positionTo int) error {
	column, findErr := mc.repository.FindById(columnId)
	if findErr != nil {
		return findErr
	}
	moveErr := mc.reorderOtherColumns(columnId, column.Position, positionTo)
	if moveErr != nil {
		return moveErr
	}
	return mc.moveColumn(columnId, positionTo)
}

func (mc *moveColumn) reorderOtherColumns(columnId int64, positionFrom int, positionTo int) error {
	if positionFrom == positionTo {
		return errors.New("COLUMN IS ALREADY IN THIS POSITION")
	}
	if positionFrom < positionTo {
		return mc.repository.MoveOtherColumnsToLeftByColumnId(columnId, positionFrom, positionTo)
	}
	return mc.repository.MoveOtherColumnsToRightByColumnId(columnId, positionFrom, positionTo)
}

func (mc *moveColumn) moveColumn(columnId int64, positionTo int) error {
	return mc.repository.Update(columnId, struct {
		Name *string
		Color *string
        Position *int
	}{ 
		Position: &positionTo, 
	})
}
